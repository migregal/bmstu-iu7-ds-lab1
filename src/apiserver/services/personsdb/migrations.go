package personsdb

import (
	"fmt"
	"log/slog"
	"sync"

	"gorm.io/gorm"

	"github.com/migregal/bmstu-iu7-ds-lab1/pkg/readiness"
)

func runMigrations(lg *slog.Logger, db *gorm.DB, probe *readiness.Probe) {
	probe.Mark(probeKey, false)

	for {
		sqlDB, err := db.DB()
		if err != nil {
			lg.Warn("[startup] failed to ping persons db", "error", err.Error())

			continue
		}

		if err = sqlDB.Ping(); err != nil {
			lg.Warn("[startup] failed to ping persons db", "error", err.Error())

			continue
		}

		break
	}

	models := map[string]any{
		"persons": Person{},
	}
	for !migrateModels(lg, db, models) { //nolint: revive
	}

	sync.OnceFunc(func() {
		probe.Mark(probeKey, true)
		lg.Warn("[startup] persons db ready")
	})()
}

func migrateModels(lg *slog.Logger, db *gorm.DB, models map[string]any) bool {
	tx := db.Begin()

	for k, v := range models {
		v := v
		if err := db.AutoMigrate(&v); err != nil {
			lg.Warn(fmt.Sprintf("[startup] failed to migrate %s", k), "err", err)
			tx.Rollback()

			return false
		}
	}

	if err := tx.Commit().Error; err != nil {
		lg.Warn("[startup] failed to commit transaction", "err", err)

		return false
	}

	return true
}
