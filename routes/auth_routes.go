package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mas-wig/post-api-1/api"
	"github.com/mas-wig/post-api-1/middleware"
	"github.com/mas-wig/post-api-1/services"
)

type AuthRoutesHandler struct{ authHandler api.AuthHandler }

func NewAuthRoutesHandler(authHandler api.AuthHandler) AuthRoutesHandler {
	return AuthRoutesHandler{authHandler: authHandler}
}

func (r *AuthRoutesHandler) AuthRouters(rg *gin.RouterGroup, userServices services.UserService) {
	router := rg.Group("auth/")
	router.POST("register", r.authHandler.SignUpUser)
	router.POST("login", r.authHandler.SignInUser)
	router.GET("refresh-token", r.authHandler.RefreshAccessToken)
	router.GET("verifyemail/:verificationCode", r.authHandler.VerifyEmail)
	router.GET("logout", middleware.DeserializeUser(userServices), r.authHandler.LogoutUser)
	router.PATCH("resetpassword/:resetToken", r.authHandler.ResetPassword)
	router.POST("forgotpassword", r.authHandler.ForgotPassword)
}
