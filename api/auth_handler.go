package api

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"text/template"

	"github.com/gin-gonic/gin"
	"github.com/mas-wig/post-api-1/config"
	"github.com/mas-wig/post-api-1/services"
	"github.com/mas-wig/post-api-1/types"
	"github.com/mas-wig/post-api-1/utils"
	"github.com/thanhpk/randstr"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthHandler struct {
	authService services.AuthService
	userService services.UserService
	ctx         context.Context
	tmpl        *template.Template
	collection  *mongo.Collection
}

func NewAuthHandle(authService services.AuthService, userService services.UserService) AuthHandler {
	return AuthHandler{authService: authService, userService: userService}
}

func (a *AuthHandler) SignUpUser(c *gin.Context) {
	var credentials *types.RegisterInput
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "Bad Request", "message": err.Error()})
		return
	}
	if credentials.Password != credentials.PasswordConfirm {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "Bad Request", "message": "password not match"})
		return
	}
	newUser, err := a.authService.RegisterUser(credentials)
	if err != nil {
		if strings.Contains(err.Error(), "email already exist") {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "Bad Request", "message": err.Error()})
			return
		}
		c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{"status": "status", "message": err.Error()})
		return
	}

	config, _ := config.LoadConfig("..")

	var (
		randomCode = randstr.String(20)
		codeVerify = utils.Encode(randomCode)
	)
	updateData := &types.UpdateInput{VerificationCode: codeVerify}
	a.userService.UpdateUserByID(newUser.ID.Hex(), updateData)
	firstName := newUser.Name

	if strings.Contains(firstName, " ") {
		firstName = strings.Split(firstName, " ")[1]
	}

	emailData := utils.EmailData{
		URL:       config.Origin + "/verifyemail/" + randomCode,
		FirstName: firstName,
		Subject:   "Your account verification code",
	}

	err = utils.SendEmail(newUser, &emailData, "verification_code.html")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{"status": "success", "message": "There was an error sending email"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"status": "success", "message": "We sent an email with a verification code to " + newUser.Email})
}

func (a *AuthHandler) SignInUser(c *gin.Context) {
	var credentials *types.LoginInput
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "Bad Request", "message": err.Error()})
		return
	}

	loginUser, err := a.userService.FindUserByEmail(credentials.Email)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "Bad Request", "message": err.Error()})
			return
		}
		return
	}

	if err := utils.VerifyPassword(loginUser.Password, credentials.Password); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	config, _ := config.LoadConfig("..")

	accessToken, err := utils.CreateToken(config.AccessTokenExpiresIn, loginUser.ID, config.AccessTokenPrivateKey)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	refreshToken, err := utils.CreateToken(config.RefreshTokenExpiresIn, loginUser.ID, config.RefreshTokenPrivateKey)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	c.SetCookie("access_token", accessToken, config.AccessTokenMaxAge*60, "/", "localhost", false, true)
	c.SetCookie("refresh_token", refreshToken, config.RefreshTokenMaxAge*60, "/", "localhost", false, true)
	c.SetCookie("logged_in", "true", config.AccessTokenMaxAge*60, "/", "localhost", false, false)
	c.JSON(http.StatusOK, gin.H{"status": "success", "access_token": accessToken})
}

func (a *AuthHandler) RefreshAccessToken(c *gin.Context) {
	cookie, err := c.Cookie("refresh_token")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "Forbidden", "message": "could not refresh access token!!"})
		return
	}

	config, _ := config.LoadConfig("..")
	sub, err := utils.ValidateToken(cookie, config.RefreshTokenPublicKey)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "Forbidden", "message": err.Error()})
		return
	}

	user, err := a.userService.FindUserByID(fmt.Sprint(sub))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "Forbidden", "message": err.Error()})
		return
	}

	accessToken, err := utils.CreateToken(config.AccessTokenExpiresIn, user.ID, config.AccessTokenPrivateKey)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "Forbidden", "message": err.Error()})
	}

	c.SetCookie("access_token", accessToken, config.AccessTokenMaxAge*60, "/", "localhost", false, true)
	c.SetCookie("logged_in", "true", config.AccessTokenMaxAge*60, "/", "localhost", false, false)
	c.JSON(http.StatusOK, gin.H{"status": "success", "access_token": accessToken})
}

func (a *AuthHandler) LogoutUser(c *gin.Context) {
	c.SetCookie("access_token", "", -1, "/", "localhost", false, true)
	c.SetCookie("refresh_token", "", -1, "/", "localhost", false, true)
	c.SetCookie("logged_in", "", -1, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
