package database

import (
	"context"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
	"github.com/zibbp/eros/ent"
	"github.com/zibbp/eros/internal/config"
)

var db *Database

type Database struct {
	Client *ent.Client
}

func InitilizeDatabase() {
	log.Debug().Msg("initializing database")

	config := config.GetConfig()

	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", config.DBHost, config.DBPort, config.DBUser, config.DBName, config.DBPassword)

	client, err := ent.Open("postgres", connectionString)

	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect to database")
	}

	// Run the auto migration tool
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatal().Err(err).Msg("failed creating schema resources")
	}

	db = &Database{
		Client: client,
	}
}

func DB() *Database {
	return db
}
