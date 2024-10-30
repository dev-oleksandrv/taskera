package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

var AppConfig *Config

func LoadConfig() (*Config, error) {
	var config Config

	env := os.Getenv("RAILWAY_ENVIRONMENT")
	if env == "" {
		env = "development"
	}

	if env == "development" {
		viper.SetConfigFile("env.development.yaml")
		viper.SetConfigType("yaml")

		if err := viper.ReadInConfig(); err != nil {
			log.Printf("Warning: Config file not found: %v", err)
		}
	}

	viper.AutomaticEnv()

	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	AppConfig = &config

	return &config, nil
}
