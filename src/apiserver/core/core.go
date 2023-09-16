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

func New(_ *slog.Logger, probe *readiness.Probe, persons persons.Client) (*Core, error) {
	return &Core{persons: persons}, nil
}

func (c *Core) AddPerson(ctx context.Context) error {
	err := c.persons.Create(ctx)
	if err != nil {
		return fmt.Errorf("failed to add person: %w", err)
	}

	return nil
}

func (c *Core) GetPerson(ctx context.Context) error {
	err := c.persons.Read(ctx)
	if err != nil {
		return fmt.Errorf("failed to add person: %w", err)
	}

	return nil
}

func (c *Core) GetPersons(ctx context.Context) error {
	err := c.persons.Read(ctx)
	if err != nil {
		return fmt.Errorf("failed to add person: %w", err)
	}

	return nil
}

func (c *Core) UpdatePerson(ctx context.Context) error {
	err := c.persons.Update(ctx)
	if err != nil {
		return fmt.Errorf("failed to add person: %w", err)
	}

	return nil
}

func (c *Core) DeletePerson(ctx context.Context) error {
	err := c.persons.Delete(ctx)
	if err != nil {
		return fmt.Errorf("failed to add person: %w", err)
	}

	return nil
}
