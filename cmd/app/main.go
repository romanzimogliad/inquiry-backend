package main

import (
	"github.com/romanzimoglyad/inquiry-backend/internal/app"
	"github.com/romanzimoglyad/inquiry-backend/internal/config"
	"github.com/romanzimoglyad/inquiry-backend/internal/logger"
	"github.com/rs/zerolog"
)

func main() {
	config.Init()
	logger.Init()
	zerolog.SetGlobalLevel(zerolog.Level(config.Config.LogLevel))
	app := app.NewApp()
	app.RunGrpcServer()
	app.RunHTTPServer()
}
