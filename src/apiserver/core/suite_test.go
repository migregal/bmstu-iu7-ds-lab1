//go:build testing
// +build testing

package core

import (
	"io"
	"log/slog"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/migregal/bmstu-iu7-ds-lab1/apiserver/core/ports/persons"
	"github.com/migregal/bmstu-iu7-ds-lab1/pkg/readiness"
)

type TestSuite struct {
	suite.Suite

	core *Core

	mockedPersons *persons.MockClient
}

func (s *TestSuite) SetupTest() {
	s.mockedPersons = persons.NewMockClient(s.T())

	var err error
	s.core, err = New(slog.New(slog.NewJSONHandler(io.Discard, nil)), readiness.New(), s.mockedPersons)

	require.NoError(s.T(), err, "failed to init core")
}

func (s *TestSuite) TearDownTest() {
}
