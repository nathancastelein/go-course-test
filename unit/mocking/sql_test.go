package mocking

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestListUsers(t *testing.T) {
	// Arrange
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error while getting mock: %s", err)
	}
	databaseService := NewDatabaseService(db)

	// TODO: use mock.ExpectQuery

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
}

func TestFindUserByIDNoRows(t *testing.T) {
}
