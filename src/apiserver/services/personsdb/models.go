package personsdb

import (
	"github.com/migregal/bmstu-iu7-ds-lab1/apiserver/core/ports/persons"
)

type Person struct {
	ID      int32 `gorm:"primaryKey"`
	Name    string
	Age     int32
	Address string
	Work    string
}

func personFromPort(p persons.Person) Person {
	return Person{
		ID:      p.ID,
		Name:    p.Name,
		Age:     p.Age,
		Address: p.Address,
		Work:    p.Work,
	}
}

func personToPort(p Person) persons.Person {
	return persons.Person{
		ID:      p.ID,
		Name:    p.Name,
		Age:     p.Age,
		Address: p.Address,
		Work:    p.Work,
	}
}
