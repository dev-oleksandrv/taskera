package config

import (
	"github.com/spf13/viper"
	"log"
)

var AppConfig *Config

func LoadConfig() (*Config, error) {
	var config Config

	viper.SetConfigFile("env.yaml")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	AppConfig = &config

	return &config, nil
}
