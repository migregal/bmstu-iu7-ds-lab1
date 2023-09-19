package v1

import "github.com/migregal/bmstu-iu7-ds-lab1/apiserver/core/ports/persons"

type PersonRequset struct {
	Name    string `json:"name" validate:"required"`
	Age     int32  `json:"age"`
	Address string `json:"address"`
	Work    string `json:"work"`
}

type PersonResponse struct {
	ID      int32  `json:"id"`
	Name    string `json:"name"`
	Age     int32  `json:"age,omitempty"`
	Address string `json:"address,omitempty"`
	Work    string `json:"work,omitempty"`
}

func personReqToPersons(p PersonRequset) persons.Person {
	return persons.Person{
		Name:    p.Name,
		Age:     p.Age,
		Address: p.Address,
		Work:    p.Work,
	}
}

func personResponseFromPersons(p persons.Person) PersonResponse {
	return PersonResponse(p)
}
