//go:build testing
// +build testing

package core

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/migregal/bmstu-iu7-ds-lab1/apiserver/core/ports/persons"
)

type PersonsSuite struct {
	TestSuite
}

func (s *PersonsSuite) SetupTest() {
	s.TestSuite.SetupTest()
}

func (s *PersonsSuite) TearDownTest() {
	s.TestSuite.TearDownTest()
}

func (s *PersonsSuite) TestAdd() {
	cases := map[string]struct {
		args  persons.Person
		dbRes int32
		dbErr error
		res   int32
		err   error
	}{
		"normal case": {
			args: persons.Person{
				Name:    "qwerty",
				Age:     18,
				Address: "asdfgh",
				Work:    "zxcvbn",
			},
		},
	}

	for tn, tc := range cases {
		s.T().Run(tn, func(t *testing.T) {
			ctx := context.Background()

			s.mockedPersons.EXPECT().Create(ctx, tc.args).Return(tc.dbRes, tc.dbErr)

			res, err := s.core.AddPerson(ctx, tc.args)
			require.Equal(s.T(), tc.res, res)
			require.ErrorIs(s.T(), err, tc.err)
		})
	}
}

func (s *PersonsSuite) TestGet() {
	cases := map[string]struct {
		id    int32
		dbRes persons.Person
		dbErr error
		res   persons.Person
		err   error
	}{
		"normal case": {
			id: 1,
			dbRes: persons.Person{
				ID:      1,
				Name:    "qwerty",
				Age:     18,
				Address: "asdfgh",
				Work:    "zxcvbn",
			},
			res: persons.Person{
				ID:      1,
				Name:    "qwerty",
				Age:     18,
				Address: "asdfgh",
				Work:    "zxcvbn",
			},
		},
		"not found case": {
			id:    2,
			dbErr: persons.ErrNotFound,
			err:   persons.ErrNotFound,
		},
	}

	for tn, tc := range cases {
		s.T().Run(tn, func(t *testing.T) {
			ctx := context.Background()

			s.mockedPersons.EXPECT().Read(ctx, tc.id).Return(tc.dbRes, tc.dbErr)

			res, err := s.core.GetPerson(ctx, tc.id)
			require.Equal(s.T(), tc.res, res)
			require.ErrorIs(s.T(), err, tc.err)
		})
	}
}

func (s *PersonsSuite) TestGetList() {
	cases := map[string]struct {
		from, to int32
		dbRes    []persons.Person
		dbErr    error
		res      []persons.Person
		err      error
	}{
		"normal case": {
			from: 1,
			to:   2,
			dbRes: []persons.Person{
				{
					ID:      1,
					Name:    "qwerty",
					Age:     18,
					Address: "asdfgh",
					Work:    "zxcvbn",
				},
			},
			res: []persons.Person{
				{
					ID:      1,
					Name:    "qwerty",
					Age:     18,
					Address: "asdfgh",
					Work:    "zxcvbn",
				},
			},
		},
		"not found case": {
			from:  2,
			to:    3,
			dbErr: persons.ErrNotFound,
			err:   persons.ErrNotFound,
		},
	}

	for tn, tc := range cases {
		s.T().Run(tn, func(t *testing.T) {
			ctx := context.Background()

			s.mockedPersons.EXPECT().ReadWithinRange(ctx, tc.from, tc.to).Return(tc.dbRes, tc.dbErr)

			res, err := s.core.GetPersons(ctx, tc.from, tc.to)
			require.Equal(s.T(), tc.res, res)
			require.ErrorIs(s.T(), err, tc.err)
		})
	}
}

func (s *PersonsSuite) TestUpdate() {
	cases := map[string]struct {
		args  persons.Person
		dbRes persons.Person
		dbErr error
		res   persons.Person
		err   error
	}{
		"normal case": {
			args: persons.Person{
				ID:   1,
				Name: "bnmghj",
				Age:  17,
			},
			dbRes: persons.Person{
				ID:      1,
				Name:    "bnmghj",
				Age:     17,
				Address: "asdfgh",
				Work:    "zxcvbn",
			},
			res: persons.Person{
				ID:      1,
				Name:    "bnmghj",
				Age:     17,
				Address: "asdfgh",
				Work:    "zxcvbn",
			},
		},
		"not found case": {
			args: persons.Person{
				ID:   2,
				Name: "qwerty",
				Age:  18,
			},
			dbErr: persons.ErrNotFound,
			err:   persons.ErrNotFound,
		},
	}

	for tn, tc := range cases {
		s.T().Run(tn, func(t *testing.T) {
			ctx := context.Background()

			s.mockedPersons.EXPECT().Update(ctx, tc.args).Return(tc.dbRes, tc.dbErr)

			res, err := s.core.UpdatePerson(ctx, tc.args)
			require.Equal(s.T(), tc.res, res)
			require.ErrorIs(s.T(), err, tc.err)
		})
	}
}

func (s *PersonsSuite) TestDelete() {
	cases := map[string]struct {
		id    int32
		dbErr error
		err   error
	}{
		"normal case": {
			id: 1,
		},
		"not found case": {
			dbErr: persons.ErrNotFound,
			err:   persons.ErrNotFound,
		},
	}

	for tn, tc := range cases {
		s.T().Run(tn, func(t *testing.T) {
			ctx := context.Background()

			s.mockedPersons.EXPECT().Delete(ctx, tc.id).Return(tc.dbErr)

			err := s.core.DeletePerson(ctx, tc.id)
			require.ErrorIs(s.T(), err, tc.err)
		})
	}
}

func TestGetPersonSuite(t *testing.T) {
	suite.Run(t, new(PersonsSuite))
}
