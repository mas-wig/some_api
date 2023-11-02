package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mas-wig/post-api-1/config"
	"github.com/mas-wig/post-api-1/services"
	"github.com/mas-wig/post-api-1/utils"
)

func DeserializeUser(userServices services.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			accessToken         string
			authorizationHeader = c.Request.Header.Get("Authorization")
			fields              = strings.Fields(authorizationHeader)
		)

		cookie, err := c.Cookie("access_token")
		if len(fields) != 0 && fields[0] == "Bearer" {
			accessToken = fields[1]
		} else if err == nil {
			accessToken = cookie
		}

		if accessToken == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "Unauthorized", "message": "you are not logged in!"})
			return
		}

		config, _ := config.LoadConfig("..")
		sub, err := utils.ValidateToken(accessToken, config.AccessTokenPublicKey)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		user, err := userServices.FindUserByID(fmt.Sprint(sub))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": err.Error()})
			return
		}
		c.Set("currentUser", user)
		c.Next()
	}
}
