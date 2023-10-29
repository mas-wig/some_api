package api

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mas-wig/post-api-1/services"
	"github.com/mas-wig/post-api-1/types"
)

type AuthHandler struct {
	authService services.AuthService
	userService services.UserService
}

func NewAuthHandle(authService services.AuthService, userService services.UserService) AuthHandler {
	return AuthHandler{authService: authService, userService: userService}
}

func (ah *AuthHandler) SignUpUser(c *gin.Context) {
	var payload *types.RegisterInput
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": err.Error()})
		return
	}
	if payload.Password != payload.PasswordConfirm {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request", "data": "password not match"})
		return
	}

	newUser, err := ah.authService.RegisterUser(payload)
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
