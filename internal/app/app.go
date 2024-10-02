package app

import (
	"effectiveMobile/internal/handlers"
	"effectiveMobile/internal/repo"

	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Run() {
	_, err := repo.ConnectToDB()
	if err != nil {
		slog.Warn("Can't connect to DB")
		return
	}
	slog.Info("Connection to Postgres [OK]")

	server := gin.Default()
	handlers.InitRouter(server)

	slog.Info("Server Started")
	http.ListenAndServe(":8080", server)
}
