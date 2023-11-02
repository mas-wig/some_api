package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mas-wig/post-api-1/services"
	"github.com/mas-wig/post-api-1/types"
)

type UserHandler struct {
	userServices services.UserService
}

func NewUserHandler(userServices services.UserService) UserHandler {
	return UserHandler{userServices: userServices}
}

func (u *UserHandler) MyProfile(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(*types.DBResponse)
	c.JSON(http.StatusOK, gin.H{"status": "success ", "message": currentUser})
}
