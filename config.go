package gobitkub

import (
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	ApiKey    string `env:"API_KEY"`
	ApiSecret string `env:"API_SECRET"`
}

func NewConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}
	config := Config{}
	if err := env.Parse(&config); err != nil {
		log.Errorf("%+v\n", err)
	}
	return config
}
