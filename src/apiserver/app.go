package apiserver

import (
	"fmt"
	"log/slog"

	"github.com/migregal/bmstu-iu7-ds-lab1/apiserver/api/http"
	"github.com/migregal/bmstu-iu7-ds-lab1/apiserver/config"
	"github.com/migregal/bmstu-iu7-ds-lab1/apiserver/core"
	"github.com/migregal/bmstu-iu7-ds-lab1/apiserver/services/personsdb"
	"github.com/migregal/bmstu-iu7-ds-lab1/pkg/apiutils"
	"github.com/migregal/bmstu-iu7-ds-lab1/pkg/readiness"
)

type App struct {
	cfg *config.Config

	http *http.Server
}

func New(lg *slog.Logger, cfg *config.Config) (*App, error) {
	a := App{cfg: cfg}

	probe := readiness.New()

	personsdb, err := personsdb.New(cfg.Persons)
	if err != nil {
		return nil, fmt.Errorf("[startup] failed to connect to personsdb: %w", err)
	}

	core, err := core.New(lg, probe, personsdb)
	if err != nil {
		return nil, fmt.Errorf("[startup] failed to init core: %w", err)
	}

	a.http, err = http.New(lg, probe, core)
	if err != nil {
		return nil, fmt.Errorf("[startup] failed to init http server: %w", err)
	}

	return &a, nil
}

func (s *App) Run(lg *slog.Logger) {
	apiutils.Serve(lg,
		apiutils.NewCallable(s.cfg.HTTPAddr, s.http),
	)
}
