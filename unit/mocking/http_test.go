package mocking

import (
	"testing"

	"github.com/h2non/gock"
)

func TestGetInfo(t *testing.T) {
	// Arrange
	defer gock.Off() // Flush pending mocks after test execution

	// TODO: add gock here!

	// Act
	info, err := GetInfo(1)

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	if info == nil {
		t.Fatal("info is not expected to be nil")
	}

	if info.FirstName != "Parisa" || info.LastName != "Tabriz" {
		t.Fatalf("expected info to be Parisa Tabriz, got %s %s", info.FirstName, info.LastName)
	}
}

func TestAddUser(t *testing.T) {
}

func TestUpdateUser(t *testing.T) {
}
