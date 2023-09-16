package personsdb

import (
	"context"
	"fmt"
	"log/slog"
	"sync"

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
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to establish connection to db: %w", err)
	}

	go pingDB(lg, db, probe)

	return &DB{db: db}, nil
}

func pingDB(lg *slog.Logger, db *gorm.DB, probe *readiness.Probe) {
	probe.Mark(probeKey, false)

	for {
		sqlDB, err := db.DB()
		if err != nil {
			lg.Warn("[startup] failed to ping persons db: %w", err)

			continue
		}

		if err = sqlDB.Ping(); err != nil {
			lg.Warn("[startup] failed to ping persons db: %w", err)

			continue
		}

		sync.OnceFunc(func() {
			probe.Mark(probeKey, true)
			lg.Warn("[startup] persons db ready")
		})()

		return
	}
}

func (d *DB) Create(ctx context.Context, p persons.Person) (int32, error) {
	return 0, nil
}

func (d *DB) Read(ctx context.Context, id int32) (persons.Person, error) {
	return persons.Person{}, nil
}

func (d *DB) Update(ctx context.Context, p persons.Person) (persons.Person, error) {
	return persons.Person{}, nil
}

func (d *DB) Delete(ctx context.Context, id int32) error {
	return nil
}
