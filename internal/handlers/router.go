package handlers

import (
	"effectiveMobile/internal/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter(server *gin.Engine) {
	server.Handle("GET", "/", Home)
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
