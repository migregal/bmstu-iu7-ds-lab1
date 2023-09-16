package config

import "github.com/migregal/bmstu-iu7-ds-lab1/apiserver/core/ports/persons"

type Config struct {
	HTTPAddr string

	Persons persons.Config
}

func ReadConfig() *Config {
	return &Config{
		HTTPAddr: ":8080",
		Persons: persons.Config{
			User:     "program",
			Password: "test",
			DBName:   "persons",
			Host:     "localhost",
			Port:     5432,
		},
	}
}
