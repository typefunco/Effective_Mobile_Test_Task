package middleware

import (
	"effectiveMobile/internal/repo"
	"effectiveMobile/internal/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Not AUTH"})
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Not AUTH"})
			c.Abort()
			return
		}

		tokenString := parts[1]
		claims, err := utils.ValidateJWT(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Not AUTH"})
			c.Abort()
			return
		}

		username := claims["username"].(string)
		isTrue, err := repo.CheckIsAdmin(username)
		if err != nil || isTrue == false {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Not Admin"})
			c.Abort()
			return
		}

		c.Next()
	}
}
