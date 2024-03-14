package mocking

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world!")
}

func TestHelloWorldHandler(t *testing.T) {
	// Arrange
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	responseRecorder := httptest.NewRecorder()

	// Act
	HelloWorldHandler(responseRecorder, request)

	// Assert
	if responseRecorder.Code != http.StatusOK {
		t.Fatalf("expected status code %d, got %d", http.StatusOK, responseRecorder.Code)
	}

	body, err := io.ReadAll(responseRecorder.Body)
	if err != nil {
		t.Fatalf("got an error while reading body: %s", err)
	}

	if string(body) != "Hello world!\n" {
		t.Fatalf("expected body Hello world!, got %s", string(body))
	}
}

func HelloHTTPCall(url string) (string, error) {
	res, err := http.Get(url + "/hello")
	if err != nil {
		return "", err
	}

	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func TestHelloHTTPCall(t *testing.T) {
	// Arrange
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("expected method GET, got %s", r.Method)
		}

		if r.URL.Path != "/hello" {
			t.Fatalf("expected path /hello, go %s", r.URL.Path)
		}

		fmt.Fprintln(w, "Hello world!")
	}))
	defer testServer.Close()

	// Act
	result, err := HelloHTTPCall(testServer.URL)

	// Assert
	if err != nil {
		t.Fatalf("got an error while making call: %s", err)
	}

	if result != "Hello world!\n" {
		t.Fatalf("expected body Hello world!, got %s", result)
	}
}
