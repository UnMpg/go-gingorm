package middleware

import (
	"fmt"
	"gin-gorm-postgres/initializers"
	"gin-gorm-postgres/models"
	"gin-gorm-postgres/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func DeserializeUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var accessToken string
		cookie, err := c.Cookie("access_token")
		authorizationHeader := c.Request.Header.Get("Authorization")
		fields := strings.Fields(authorizationHeader)

		if len(fields) != 0 && fields[0] == "Bearer" {
			accessToken = fields[1]
		} else if err == nil {
			accessToken = cookie
		}

		if accessToken == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "You not logged in"})
			return
		}
		config, _ := initializers.LoadConfig(".")
		sub, err := utils.ValidateToken(accessToken, config.AccessTokenPublicKey)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		var user models.User
		result := initializers.DB.First(&user, "id= ?", fmt.Sprint(sub))

		if result.Error != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": "the user belong to this token no logger exist"})
			return
		}

		c.Set("currentUser", user)
		c.Next()
	}
}
