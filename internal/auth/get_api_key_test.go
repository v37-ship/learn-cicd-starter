package auth

import (
	"net/http"
	"testing"
)

func TestGetApiKey(t *testing.T) {
	// 1. Test a valid header
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey secret-key")

	got, err := GetAPIKey(headers)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	if got != "secret-key" {
		t.Errorf("got %q, want %q", got, "secret-key")
	}

	// 2. Test an empty/missing header
	emptyHeaders := http.Header{}
	_, err = GetAPIKey(emptyHeaders)
	if err == nil {
		t.Errorf("expected an error, but got nil")
	}

	// 3. Test an wrong format header
	wrongHeaders := http.Header{}
	wrongHeaders.Set("Authorization", "ApiKey")

	_, formatErr := GetAPIKey(wrongHeaders)
	if formatErr == nil {
		t.Errorf("expected an error, but got nil")
	}
}
