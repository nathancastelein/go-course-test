package mocking

import (
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/nathancastelein/go-course-test/unit/mocking"
)

func TestListUsers(t *testing.T) {
	// Arrange
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error while getting mock: %s", err)
	}
	databaseService := mocking.NewDatabaseService(db)

	mock.ExpectQuery("SELECT id, first_name, last_name FROM users").
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "first_name", "last_name"}).
				AddRow(1, "Grace", "Hopper"),
		)

	// Act
	users, err := databaseService.ListUsers()

	// Assert
	if err != nil {
		t.Fatalf("error while listing users: %s", err)
	}

	if len(users) != 1 {
		t.Fatalf("expected users length %d, got %d", 1, len(users))
	}

	user := users[0]
	if user.ID != 1 || user.FirstName != "Grace" || user.LastName != "Hopper" {
		t.Fatal("invalid user received from database")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("mock expectations failed: %s", err)
	}
}

func TestFindUserByID(t *testing.T) {
	// Arrange
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error while getting mock: %s", err)
	}
	databaseService := mocking.NewDatabaseService(db)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT id, first_name, last_name FROM users WHERE id = $1")).
		WithArgs(1).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "first_name", "last_name"}).
				AddRow(1, "Grace", "Hopper"),
		)

	// Act
	user, err := databaseService.FindUserByID(1)

	// Assert
	if err != nil {
		t.Fatalf("error while finding user: %s", err)
	}

	if user == nil {
		t.Fatal("expected user not nil")
	}

	if user.ID != 1 || user.FirstName != "Grace" || user.LastName != "Hopper" {
		t.Fatal("invalid user received from database")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("mock expectations failed: %s", err)
	}
}

func TestFindUserByIDNoRows(t *testing.T) {
	// Arrange
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error while getting mock: %s", err)
	}
	databaseService := mocking.NewDatabaseService(db)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT id, first_name, last_name FROM users WHERE id = $1")).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "first_name", "last_name"}))

	// Act
	user, err := databaseService.FindUserByID(1)

	// Assert
	if err != sql.ErrNoRows {
		t.Fatalf("expecting error sql.ErrNoRows, got: %s", err)
	}

	if user != nil {
		t.Fatal("expecting user to be nil")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("mock expectations failed: %s", err)
	}
}
