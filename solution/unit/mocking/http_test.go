package mocking

import (
	"net/http"
	"testing"

	"github.com/h2non/gock"
	"github.com/nathancastelein/go-course-test/unit/mocking"
)

func TestGetInfo(t *testing.T) {
	// Arrange
	defer gock.Off() // Flush pending mocks after test execution

	gock.New("https://yourapitomock.io").
		Get("/").
		MatchParam("user", "1").
		Reply(http.StatusOK).
		JSON(map[string]string{
			"first_name": "Parisa",
			"last_name":  "Tabriz",
		})

	// Act
	info, err := mocking.GetInfo(1)

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
	// Arrange
	defer gock.Off() // Flush pending mocks after test execution

	gock.New("https://yourapitomock.io").
		Post("/user").
		MatchType("application/json").
		JSON(&mocking.Information{
			FirstName: "Grace",
			LastName:  "Hopper",
		}).
		Reply(http.StatusOK).
		JSON(map[string]int{
			"id": 42,
		})

	// Act
	id, err := mocking.AddUser(&mocking.Information{
		FirstName: "Grace",
		LastName:  "Hopper",
	})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	if id == 0 {
		t.Fatal("id is not expected to be 0")
	}
}

func TestUpdateUser(t *testing.T) {
	// Arrange
	defer gock.Off() // Flush pending mocks after test execution

	gock.New("https://yourapitomock.io").
		Put("/user/1").
		MatchType("application/json").
		JSON(&mocking.Information{
			FirstName: "Grace",
			LastName:  "Hopper",
		}).
		Reply(http.StatusNoContent)

	// Act
	err := mocking.UpdateUser(1, &mocking.Information{
		FirstName: "Grace",
		LastName:  "Hopper",
	})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
}
