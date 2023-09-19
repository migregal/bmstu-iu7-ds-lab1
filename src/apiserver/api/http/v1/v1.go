package v1

import (
	"context"
	"errors"
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
	gr.GET("/persons/:id", a.GetPerson)
	gr.PATCH("/persons/:id", a.PatchPerson)
	gr.DELETE("/persons/:id", a.DeletePerson)

	return nil
}

type api struct {
	core Core
}

func (a *api) PostPerson(c echo.Context) error {
	var (
		req PersonRequset
		err error
	)

	if err = c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = c.Validate(req); err != nil {
		return err //nolint: wrapcheck
	}

	id, err := a.core.AddPerson(c.Request().Context(), personReqToPersons(req))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	c.Response().Header().Add("Location", fmt.Sprintf("/api/v1/persons/%d", id))

	return c.NoContent(http.StatusCreated)
}

func (a *api) GetPersons(c echo.Context) error {
	ps, err := a.core.GetPersons(c.Request().Context(), 0, math.MaxInt32)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	infos := make([]PersonResponse, 0, len(ps))
	for _, p := range ps {
		infos = append(infos, personResponseFromPersons(p))
	}

	return c.JSON(http.StatusOK, infos)
}

func (a *api) GetPerson(c echo.Context) error {
	id64, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	id := int32(id64)

	p, err := a.core.GetPerson(c.Request().Context(), id)
	if err != nil {
		if errors.Is(err, persons.ErrNotFound) {
			return c.NoContent(http.StatusNotFound)
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, personResponseFromPersons(p))
}

func (a *api) PatchPerson(c echo.Context) error {
	id64, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	var req PersonRequset
	if err = c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = c.Validate(req); err != nil {
		return err //nolint: wrapcheck
	}

	p := personReqToPersons(req)
	p.ID = int32(id64)

	p, err = a.core.UpdatePerson(c.Request().Context(), p)
	if err != nil {
		if errors.Is(err, persons.ErrNotFound) {
			return c.NoContent(http.StatusNotFound)
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, personResponseFromPersons(p))
}

func (a *api) DeletePerson(c echo.Context) error {
	id64, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	err = a.core.DeletePerson(c.Request().Context(), int32(id64))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}
