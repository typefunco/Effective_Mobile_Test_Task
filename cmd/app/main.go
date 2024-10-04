package main

import (
	"effectiveMobile/internal/app"
	"time"
)

// @title Effective Mobile API
// @version 1.0
// @description This is the API for Effective Mobile project
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization

func main() {
	// wait db
	time.Sleep(time.Second * 10)
	app.Run()
}
