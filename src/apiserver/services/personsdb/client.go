package personsdb

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/migregal/bmstu-iu7-ds-lab1/apiserver/core/ports/persons"
	"github.com/migregal/bmstu-iu7-ds-lab1/pkg/readiness"
)

const probeKey = "personsdb"

type DB struct {
	db *gorm.DB
}

func New(lg *slog.Logger, cfg persons.Config, probe *readiness.Probe) (*DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to establish connection to db: %w", err)
	}

	go runMigrations(lg, db, probe)

	return &DB{db: db}, nil
}

func (d *DB) Create(_ context.Context, p persons.Person) (int32, error) {
	data := personFromPort(p)

	if err := d.db.Create(&data).Error; err != nil {
		return 0, fmt.Errorf("failed to create person: %w", err)
	}

	return data.ID, nil
}

func (d *DB) Read(_ context.Context, id int32) (persons.Person, error) {
	p := Person{ID: id}

	if err := d.db.First(&p).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = persons.ErrNotFound
		}

		return persons.Person{}, fmt.Errorf("failed to read persons: %w", err)
	}

	return personToPort(p), nil
}

func (d *DB) ReadWithinRange(_ context.Context, from, to int32) ([]persons.Person, error) {
	ps := []Person{}

	if err := d.db.Limit(int(to - from)).Offset(int(from)).Find(&ps).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = persons.ErrNotFound
		}

		return nil, fmt.Errorf("failed to read persons: %w", err)
	}

	res := make([]persons.Person, 0, len(ps))
	for _, p := range ps {
		res = append(res, personToPort(p))
	}

	return res, nil
}

func (d *DB) Update(_ context.Context, p persons.Person) (persons.Person, error) {
	data := personFromPort(p)

	if err := d.db.Save(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return persons.Person{}, persons.ErrNotFound
		}

		return persons.Person{}, fmt.Errorf("failed to update person: %w", err)
	}

	return persons.Person{}, nil
}

func (d *DB) Delete(_ context.Context, id int32) error {
	p := Person{ID: id}

	if err := d.db.Delete(&p).Error; err != nil {
		return fmt.Errorf("failed to delete person: %w", err)
	}

	return nil
}
