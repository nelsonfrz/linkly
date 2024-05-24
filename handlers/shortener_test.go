package handlers_test

import (
	"linkly/handlers"
	"linkly/storage"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestShortenHandler(t *testing.T) {
	// Create a request to pass to our handler
	req, err := http.NewRequest("GET", "/shorten?url=https://example.com", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.ShortenHandler)

	// Serve the request
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect
	expected := handlers.BaseUrl + "/r/1"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestRedirectHandler(t *testing.T) {
	// Save a URL to setup the test
	id := storage.SaveUrl("https://example.com")

	// Create a request to pass to our handler
	req, err := http.NewRequest("GET", "/r/"+id, nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// We need to create a router and register our handler with it
	router := mux.NewRouter()
	router.HandleFunc("/r/{id}", handlers.RedirectHandler)

	// Serve the request
	router.ServeHTTP(rr, req)

	// Check the status code is what we expect
	if status := rr.Code; status != http.StatusFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusFound)
	}

	// Check the Location header is what we expect
	expected := "https://example.com"
	if location := rr.Header().Get("Location"); location != expected {
		t.Errorf("handler returned wrong location header: got %v want %v",
			location, expected)
	}
}
