package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
	"github.com/zibbp/eros/internal/config"
	"github.com/zibbp/eros/internal/database"
	"github.com/zibbp/eros/internal/report"
	"github.com/zibbp/eros/internal/s3"
	"github.com/zibbp/eros/internal/script"
	transportHttp "github.com/zibbp/eros/internal/transport/http"
)

func Run() error {

	if os.Getenv("ENV") == "development" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal().Err(err).Msg("Error loading .env file")
		}
	}

	config := config.GetConfig()

	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	if config.Debug {
		log.Info().Msg("debug mode enabled")
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	// init database
	database.InitilizeDatabase()
	db := database.DB()

	// init s3
	s3, err := s3.NewS3Client(config.S3_ENDPOINT, config.S3_ACCESSKEYID, config.S3_ACCESSKEYSECRET, config.S3_BUCKET)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to init s3 client")
	}

	log.Debug().Msg("initilized s3 client")

	// init services

	scriptService := script.NewService(db)
	reportService := report.NewService(db, s3, scriptService)

	httpHandler := transportHttp.NewHandler(scriptService, reportService)

	if err := httpHandler.Serve(); err != nil {
		return err
	}

	return nil
}

func main() {
	if err := Run(); err != nil {
		log.Fatal().Err(err).Msg("failed to run")
	}
}
