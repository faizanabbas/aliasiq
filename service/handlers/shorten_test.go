package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestShortenURL(t *testing.T) {
	reqBody := bytes.NewBuffer([]byte(`{"originalUrl": "https://localhost:8081"}`))
	req, err := http.NewRequest("POST", "/shorten", reqBody)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	ShortenURL(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status code %d, got %d", 200, w.Code)
	}

	var response map[string]string
	err = json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}
	if response["shortUrl"] != "https://localhost:8080/redirect/"+"alias123" {
		t.Errorf(
			"expected shortUrl to be %s, got %s",
			"https://localhost:8080/redirect/"+"alias123",
			response["shortUrl"],
		)
	}
}
