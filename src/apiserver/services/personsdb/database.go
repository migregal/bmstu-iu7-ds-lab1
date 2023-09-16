package personsdb

import (
	"context"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/migregal/bmstu-iu7-ds-lab1/apiserver/core/ports/persons"
)

type DB struct {
	db *gorm.DB
}

func New(cfg persons.Config) (*DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to establish connection to db: %w", err)
	}

	return &DB{db: db}, nil
}

func (d *DB) Create(ctx context.Context) error {
	return nil
}

func (d *DB) Read(ctx context.Context) error {
	return nil
}

func (d *DB) Update(ctx context.Context) error {
	return nil
}

func (d *DB) Delete(ctx context.Context) error {
	return nil
}
