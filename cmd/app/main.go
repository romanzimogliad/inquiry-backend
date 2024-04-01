package main

import (
	"context"
	"fmt"

	"github.com/romanzimoglyad/inquiry-backend/internal/domain"

	"github.com/romanzimoglyad/inquiry-backend/internal/database"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/romanzimoglyad/inquiry-backend/internal/app"
	"github.com/romanzimoglyad/inquiry-backend/internal/config"
	"github.com/romanzimoglyad/inquiry-backend/internal/logger"
	"github.com/rs/zerolog"
)

func main() {
	//ctx := context.Background()
	config.Init()
	logger.Init()
	zerolog.SetGlobalLevel(zerolog.Level(config.Config.LogLevel))

	psqlConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Config.Database.Host,
		config.Config.Database.Port,
		config.Config.Database.User,
		config.Config.Database.Password,
		config.Config.Database.Name,
	)

	dbpool, err := pgxpool.Connect(context.Background(), psqlConn)
	if err != nil {
		logger.Fatal().Msgf("Unable to create connection pool: %v\n", err)
	}
	defer dbpool.Close()

	database := database.New(dbpool)
	inquiryService := domain.New(database)

	app := app.NewApp(inquiryService)
	app.RunGrpcServer()
	app.RunHTTPServer()

}
