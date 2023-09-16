package core

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/migregal/bmstu-iu7-ds-lab1/apiserver/core/ports/persons"
	"github.com/migregal/bmstu-iu7-ds-lab1/pkg/readiness"
)

type Core struct {
	persons persons.Client
}

func New(lg *slog.Logger, probe *readiness.Probe, persons persons.Client) (*Core, error) {
	probe.Mark("core", true)
	lg.Warn("[startup] core ready")

	return &Core{persons: persons}, nil
}

func (c *Core) AddPerson(ctx context.Context, p persons.Person) (int32, error) {
	id, err := c.persons.Create(ctx, p)
	if err != nil {
		return 0, fmt.Errorf("failed to add person: %w", err)
	}

	return id, nil
}

func (c *Core) GetPerson(ctx context.Context, id int32) (persons.Person, error) {
	p, err := c.persons.Read(ctx, id)
	if err != nil {
		return persons.Person{}, fmt.Errorf("failed to add person: %w", err)
	}

	return p, nil
}

func (c *Core) GetPersons(ctx context.Context, from, to int32) ([]persons.Person, error) {
	_, err := c.persons.Read(ctx, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to add person: %w", err)
	}

	return nil, nil
}

func (c *Core) UpdatePerson(ctx context.Context, p persons.Person) (persons.Person, error) {
	p, err := c.persons.Update(ctx, p)
	if err != nil {
		return persons.Person{}, fmt.Errorf("failed to add person: %w", err)
	}

	return p, nil
}

func (c *Core) DeletePerson(ctx context.Context, id int32) error {
	err := c.persons.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to add person: %w", err)
	}

	return nil
}
