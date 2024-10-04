package handlers

import (
	"effectiveMobile/internal/middleware"

	_ "effectiveMobile/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter(server *gin.Engine) {
	server.Handle("GET", "/", Home)
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	server.Handle("POST", "/sign-up", signUp)
	server.Handle("PATCH", "/update", becomeAdmin)

	protected := server.Group("/music")
	protected.Use(middleware.JWTMiddleware())
	{
		protected.Handle("GET", "/songs", GetAllSongs)
		protected.Handle("GET", "/songs/:song_id", GetSongs)
		protected.Handle("GET", "/song/:song_id/:verse", GetSong)
	}

	admin := server.Group("/music")
	admin.Use(middleware.AdminMiddleware())
	{
		admin.Handle("DELETE", "/song/:id", DeleteSong)
		admin.Handle("PATCH", "/songs/:song_id", UpdateSong)
		admin.Handle("POST", "/song/new", NewSong)
	}
}
