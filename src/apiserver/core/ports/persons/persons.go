package persons

import (
	"context"
	"fmt"
)

var (
	ErrNotFound = fmt.Errorf("not found")
)

type Config struct {
	User     string
	Password string
	DBName   string
	Host     string
	Port     int
}

type Client interface {
	Create(ctx context.Context) error
	Read(ctx context.Context) error
	Update(ctx context.Context) error
	Delete(ctx context.Context) error
}
