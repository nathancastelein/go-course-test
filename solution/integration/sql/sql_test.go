//go:build integration
// +build integration

package sql

import (
	"context"
	"database/sql"
	"path/filepath"
	"testing"
	"time"

	_ "github.com/lib/pq"
	"github.com/nathancastelein/go-course-test/unit/mocking"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

type DatabaseTestSuite struct {
	suite.Suite
	context   context.Context
	db        *sql.DB
	container *postgres.PostgresContainer
	service   *mocking.DatabaseService
}

func (t *DatabaseTestSuite) SetupSuite() {
	t.context = context.Background()

	postgresContainer, err := postgres.RunContainer(t.context,
		testcontainers.WithImage("docker.io/postgres:alpine"),
		postgres.WithInitScripts(filepath.Join("testdata", "init.sql")),
		postgres.WithDatabase("test"),
		postgres.WithUsername("user"),
		postgres.WithPassword("password"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second)),
	)

	t.Require().NoError(err)
	t.container = postgresContainer

	connStr, err := postgresContainer.ConnectionString(t.context, "sslmode=disable", "application_name=test")
	t.Require().NoError(err)

	db, err := sql.Open("postgres", connStr)
	t.Require().NoError(err)

	t.db = db
	t.service = mocking.NewDatabaseService(t.db)
}

func (t *DatabaseTestSuite) TearDownSuite() {
	t.Require().NoError(t.container.Terminate(t.context))
}

func (t *DatabaseTestSuite) TestListUsers() {
	// Act
	users, err := t.service.ListUsers()

	// Assert
	t.Require().NoError(err)
	t.Require().Len(users, 2, "length should be two")
}

func TestDatabaseService(t *testing.T) {
	suite.Run(t, &DatabaseTestSuite{})
}
