package config

import (
	"fmt"

	"github.com/spf13/viper"

	"github.com/migregal/bmstu-iu7-ds-lab1/apiserver/core/ports/persons"
)

type Config struct {
	HTTPAddr string `mapstructure:"http_addr"`

	Persons persons.Config
}

func ReadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/etc/apiserver/")
	viper.AddConfigPath(".")

	viper.SetDefault("http_addr", ":8080")
	viper.SetDefault("persons.user", "program")
	viper.SetDefault("persons.password", "test")
	viper.SetDefault("persons.database", "persons")
	viper.SetDefault("persons.host", "localhost")
	viper.SetDefault("persons.port", "5432")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("[startup] failed to read config: %w", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("[startup] failed to parse config: %w", err)
	}

	return &cfg, nil
}
