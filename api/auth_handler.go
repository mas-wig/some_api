package api

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mas-wig/post-api-1/config"
	"github.com/mas-wig/post-api-1/services"
	"github.com/mas-wig/post-api-1/types"
	"github.com/mas-wig/post-api-1/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthHandler struct {
	authService services.AuthService
	userService services.UserService
}

func NewAuthHandle(authService services.AuthService, userService services.UserService) AuthHandler {
	return AuthHandler{authService: authService, userService: userService}
}

func (a *AuthHandler) SignUpUser(c *gin.Context) {
	var credentials *types.RegisterInput
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": err.Error()})
		return
	}
	if credentials.Password != credentials.PasswordConfirm {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request", "data": "password not match"})
		return
	}

	newUser, err := a.authService.RegisterUser(credentials)
	if err != nil {
		if strings.Contains(err.Error(), "email already exist") {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request", "data": err.Error()})
			return
		}
		c.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "success", "data": newUser})
}

func (a *AuthHandler) SignInUser(c *gin.Context) {
	var credentials *types.LoginInput
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": err.Error()})
		return
	}

	loginUser, err := a.userService.FindUserByEmail(credentials.Email)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request", "data": err.Error()})
			return
		}
		return
	}

	if err := utils.VerifyPassword(loginUser.Password, credentials.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request", "data": err.Error()})
		return
	}

	config, _ := config.LoadConfig("..")

	accessToken, err := utils.CreateToken(config.AccessTokenExpiresIn, loginUser.ID, config.AccessTokenPrivateKey)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request", "data": err.Error()})
		return
	}
	refreshToken, err := utils.CreateToken(config.RefreshTokenExpiresIn, loginUser.ID, config.RefreshTokenPrivateKey)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	c.SetCookie("access_token", accessToken, config.AccessTokenMaxAge*60, "/", "localhost", false, true)
	c.SetCookie("refresh_token", refreshToken, config.RefreshTokenMaxAge*60, "/", "localhost", false, true)
	c.SetCookie("logged_in", "true", config.AccessTokenMaxAge*60, "/", "localhost", false, false)
	c.JSON(http.StatusOK, gin.H{"status": "success", "access_token": accessToken})
}
