package handlers

import (
	"effectiveMobile/internal/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter(server *gin.Engine) {
	server.Handle("GET", "/", Home)
	server.Handle("POST", "/sign-up", signUp)
	server.Handle("PATCH", "/update", becomeAdmin)

	protected := server.Group("/protected")
	protected.Use(middleware.JWTMiddleware())
	{
		// hadnlers
	}
}
