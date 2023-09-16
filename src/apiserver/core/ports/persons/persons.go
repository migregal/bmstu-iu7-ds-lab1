package persons

import (
	"context"
	"fmt"
)

var ErrNotFound = fmt.Errorf("not found")

type Config struct {
	User     string
	Password string
	DBName   string
	Host     string
	Port     int
}

type Client interface {
	Create(context.Context, Person) (int32, error)
	Read(context.Context, int32) (Person, error)
	ReadWithinRange(context.Context, int32, int32) ([]Person, error)
	Update(context.Context, Person) (Person, error)
	Delete(context.Context, int32) error
}

type Person struct {
	ID      int32
	Name    string
	Age     int32
	Address string
	Work    string
}
