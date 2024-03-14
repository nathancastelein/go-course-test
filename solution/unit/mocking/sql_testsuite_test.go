package mocking

import (
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/nathancastelein/go-course-test/unit/mocking"
	"github.com/stretchr/testify/suite"
)

type TestDatabaseServiceSuite struct {
	suite.Suite

	database     *sql.DB
	databaseMock sqlmock.Sqlmock
	service      *mocking.DatabaseService
}

func (t *TestDatabaseServiceSuite) SetupSuite() {
	var err error
	t.database, t.databaseMock, err = sqlmock.New()

	t.Require().NoError(err)

	t.service = mocking.NewDatabaseService(t.database)
}

func (t *TestDatabaseServiceSuite) TearDownSuite() {}

func (t *TestDatabaseServiceSuite) SetupTest() {}

func (t *TestDatabaseServiceSuite) TearDownTest() {
	t.Require().NoError(t.databaseMock.ExpectationsWereMet())
}

func (t *TestDatabaseServiceSuite) TestListUsers() {
	t.databaseMock.ExpectQuery("SELECT id, first_name, last_name FROM users").
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "first_name", "last_name"}).
				AddRow(1, "Grace", "Hopper"),
		)

	// Act
	users, err := t.service.ListUsers()

	// Assert
	t.Require().NoError(err)
	t.Require().Len(users, 1, "length should be one")
	t.Require().Equal(users[0], &mocking.User{ID: 1, FirstName: "Grace", LastName: "Hopper"}, "they should be equal")
}

func (t *TestDatabaseServiceSuite) TestFindUserByID() {
	t.databaseMock.ExpectQuery(regexp.QuoteMeta("SELECT id, first_name, last_name FROM users WHERE id = $1")).
		WithArgs(1).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "first_name", "last_name"}).
				AddRow(1, "Grace", "Hopper"),
		)

	// Act
	user, err := t.service.FindUserByID(1)

	// Assert
	t.Require().NoError(err)
	t.Require().NotNil(user)
	t.Require().Equal(user, &mocking.User{ID: 1, FirstName: "Grace", LastName: "Hopper"}, "they should be equal")
}

func (t *TestDatabaseServiceSuite) TestFindUserByIDNoRows() {
	t.databaseMock.ExpectQuery(regexp.QuoteMeta("SELECT id, first_name, last_name FROM users WHERE id = $1")).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "first_name", "last_name"}))

	// Act
	user, err := t.service.FindUserByID(1)

	// Assert
	t.Require().ErrorIs(err, sql.ErrNoRows)
	t.Require().Nil(user)
}

func TestDatabaseService(t *testing.T) {
	suite.Run(t, &TestDatabaseServiceSuite{})
}
