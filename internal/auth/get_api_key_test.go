package auth

import (
	"errors"
	"net/http"
	"testing"
)

// REMOVED: Do not redeclare ErrNoAuthHeaderIncluded here if it exists in your source file.

func TestGetAPIKey_Valid(t *testing.T) {
	headers := http.Header{
		"Authorization": []string{"ApiKey secret123"},
	}

	key, err := GetAPIKey(headers)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if key != "secret123" {
		t.Errorf("expected key 'secret123', got '%s'", key)
	}
}

func TestGetAPIKey_MissingHeader(t *testing.T) {
	headers := http.Header{}

	_, err := GetAPIKey(headers)
	if !errors.Is(err, ErrNoAuthHeaderIncluded) {
		t.Errorf("expected ErrNoAuthHeaderIncluded, got %v", err)
	}
}
