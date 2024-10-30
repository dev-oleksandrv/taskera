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
		viper.SetConfigFile(".env.development")
		viper.SetConfigType("dotenv")

		if err := viper.ReadInConfig(); err != nil {
			log.Printf("Warning: Config file not found: %v", err)
		}
	}

	viper.AutomaticEnv()

	config.Server.Port = viper.GetString("SERVER_PORT")
	config.Server.Env = viper.GetString("SERVER_ENV")
	config.Database.Host = viper.GetString("DATABASE_HOST")
	config.Database.Port = viper.GetInt("DATABASE_PORT")
	config.Database.Username = viper.GetString("DATABASE_USERNAME")
	config.Database.Password = viper.GetString("DATABASE_PASSWORD")
	config.Database.Name = viper.GetString("DATABASE_NAME")
	config.Auth.Secret = viper.GetString("AUTH_SECRET")

	AppConfig = &config

	return &config, nil
}
