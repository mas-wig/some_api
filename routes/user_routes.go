package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mas-wig/post-api-1/api"
	"github.com/mas-wig/post-api-1/middleware"
	"github.com/mas-wig/post-api-1/services"
)

type UserRoutesHandler struct{ userHandler api.UserHandler }

func NewUserRoutesHandler(userHandler api.UserHandler) UserRoutesHandler {
	return UserRoutesHandler{userHandler: userHandler}
}

func (r *UserRoutesHandler) UserRouters(rg *gin.RouterGroup, userServices services.UserService) {
	router := rg.Group("users/")
	router.Use(middleware.DeserializeUser(userServices))
	router.GET("/me", r.userHandler.MyProfile)
}
