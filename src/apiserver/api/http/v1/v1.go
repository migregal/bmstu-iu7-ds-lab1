package v1

import (
	"context"
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/migregal/bmstu-iu7-ds-lab1/apiserver/core/ports/persons"
)

type Core interface {
	AddPerson(context.Context, persons.Person) (int32, error)
	GetPerson(context.Context, int32) (persons.Person, error)
	GetPersons(context.Context, int32, int32) ([]persons.Person, error)
	UpdatePerson(context.Context, persons.Person) (persons.Person, error)
	DeletePerson(context.Context, int32) error
}

func InitListener(mx *echo.Echo, core Core) error {
	gr := mx.Group("/api/v1")

	a := api{core: core}

	gr.POST("/persons", a.PostPerson)
	gr.GET("/persons", a.GetPersons)
	gr.GET("/persons/:id", a.GetPersons)
	gr.PATCH("/persons/:id", a.PatchPerson)
	gr.DELETE("/persons/:id", a.DeletePerson)

	return nil
}

type api struct {
	core Core
}

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

type ValidationErrorResponse struct {
	Message string `json:"message"`
	Errors  string `json:"errors"`
}

func (a *api) PostPerson(c echo.Context) error {
	var (
		req PersonRequset
		err error
	)

	if err = c.Bind(&req); err != nil {
		// return c.JSON(http.StatusBadRequest, ValidationErrorResponse{})
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	_, err = a.core.AddPerson(c.Request().Context(), persons.Person{})
	if err != nil {
		return fmt.Errorf("failed to add new person: %w", err)
	}

	info := PersonResponse{}

	return c.JSON(http.StatusCreated, info)
}

func (a *api) GetPersons(c echo.Context) error {
	infos := make([]PersonResponse, 1)

	_, err := a.core.GetPersons(c.Request().Context(), 0, math.MaxInt32)
	if err != nil {
		return fmt.Errorf("failed to get list of persons: %w", err)
	}

	return c.JSON(http.StatusOK, infos)
}

func (a *api) GetPerson(c echo.Context) error {
	id64, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	id := int32(id64)

	_, err = a.core.GetPerson(c.Request().Context(), id)
	if err != nil {
		return fmt.Errorf("failed to get person: %w", err)
	}

	info := PersonResponse{ID: id}

	return c.JSON(http.StatusOK, info)
}

func (a *api) PatchPerson(c echo.Context) error {
	id64, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	id := int32(id64)

	var req PersonRequset
	if err = c.Bind(&req); err != nil {
		// return c.JSON(http.StatusBadRequest, ValidationErrorResponse{})
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	_, err = a.core.UpdatePerson(c.Request().Context(), persons.Person{})
	if err != nil {
		return fmt.Errorf("failed to update person: %w", err)
	}

	info := PersonResponse{ID: id}

	return c.JSON(http.StatusOK, info)
}

func (a *api) DeletePerson(c echo.Context) error {
	id64, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	id := int32(id64)

	var req PersonRequset
	if err = c.Bind(&req); err != nil {
		// return c.JSON(http.StatusBadRequest, ValidationErrorResponse{})
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = a.core.DeletePerson(c.Request().Context(), id)
	if err != nil {
		return fmt.Errorf("failed to delete person: %w", err)
	}

	info := PersonResponse{ID: id}

	return c.JSON(http.StatusOK, info)
}
