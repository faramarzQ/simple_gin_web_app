package middlewares

import (
	"gin/helpers"

	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		if token, err := c.Cookie("token"); err == nil || token != "" {
			c.Set("is_logged_in", true)
			c.Next()
		} else {
			helpers.LogDanger("Unauthorized request!")
			c.Set("is_logged_in", false)
			c.Abort()
		}
	}
}
