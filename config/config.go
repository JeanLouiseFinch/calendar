package config

import (
	"errors"

	"github.com/spf13/viper"
)

type Config struct {
	IP      string
	Port    int
	LogFile string
	Detail  string
}

func GetConfig(file string) (*Config, error) {
	viper.AutomaticEnv()
	viper.SetConfigFile(file)

	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.New("Error reading config")
	}
	var config Config

	err := viper.Unmarshal(&config)
	if err != nil {
		return nil, errors.New("Error unmarshal config")
	}
	return &config, nil
}
