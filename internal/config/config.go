package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog/log"
)

type Config struct {
	Debug              bool   `envconfig:"DEBUG" default:"false"`
	DBHost             string `envconfig:"DB_HOST" default:"localhost"`
	DBName             string `envconfig:"DB_NAME" default:"eros"`
	DBUser             string `envconfig:"DB_USER" default:"eros"`
	DBPassword         string `envconfig:"DB_PASSWORD" default:"eros"`
	DBPort             string `envconfig:"DB_PORT" default:"5432"`
	S3_ENDPOINT        string `envconfig:"S3_ENDPOINT" default:"http://localhost:9000"`
	S3_ACCESSKEYID     string `envconfig:"S3_ACCESS_KEY_ID" default:"minio"`
	S3_ACCESSKEYSECRET string `envconfig:"S3_ACCESS_KEY_SECRET" default:"minio123"`
	S3_BUCKET          string `envconfig:"S3_BUCKET" default:"eros"`
}

func GetConfig() *Config {
	var c Config
	if err := envconfig.Process("", &c); err != nil {
		log.Fatal().Err(err).Msg("failed to load config")
	}

	return &c
}
