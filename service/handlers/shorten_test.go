package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"
)

func TestShortenURL(t *testing.T) {
	req, err := http.NewRequest(
		"POST",
		"/shorten",
		bytes.NewBuffer([]byte(`{"originalUrl": "https://localhost:8081"}`)),
	)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	ShortenURL(w, req)

	expectedStatusCode := http.StatusOK
	if w.Code != expectedStatusCode {
		t.Errorf("expected status code %d, got %d", expectedStatusCode, w.Code)
	}

	var response map[string]string
	err = json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	shortURL := response["shortUrl"]
	baseURL := "https://localhost:8080/redirect/"
	if !strings.HasPrefix(shortURL, baseURL) {
		t.Errorf("expected shortUrl to start with %s, got %s", baseURL, shortURL)
	}

	alias := strings.TrimPrefix(shortURL, baseURL)
	expectedAliasLength := 7
	if len(alias) != expectedAliasLength {
		t.Errorf("expected alias length to be %d, got %d", expectedAliasLength, len(alias))
	}
	if !regexp.MustCompile(`^[a-z0-9]+$`).MatchString(alias) {
		t.Errorf("expected alias to match [a-z0-9]+, got %s", alias)
	}
}

func TestShortenURL_MissingOriginalURL(t *testing.T) {
	req, err := http.NewRequest(
		"POST",
		"/shorten",
		bytes.NewBuffer([]byte(`{}`)),
	)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	ShortenURL(w, req)

	expectedStatusCode := http.StatusBadRequest
	if w.Code != expectedStatusCode {
		t.Errorf("expected status code %d, got %d", expectedStatusCode, w.Code)
	}

	var response map[string]string
	err = json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	expectedError := "missing or invalid originalUrl"
	if response["error"] != expectedError {
		t.Errorf("expected error message '%s', got %s", expectedError, response["error"])
	}
}

func TestShortenURL_InvalidOriginalURL(t *testing.T) {
	req, err := http.NewRequest(
		"POST",
		"/shorten",
		bytes.NewBuffer([]byte(`{"originalUrl": "not-a-valid-url"}`)),
	)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	ShortenURL(w, req)

	expectedStatusCode := http.StatusBadRequest
	if w.Code != expectedStatusCode {
		t.Errorf("expected status code %d, got %d", expectedStatusCode, w.Code)
	}

	var response map[string]string
	err = json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	expectedError := "invalid URL format"
	if response["error"] != expectedError {
		t.Errorf("expected error message '%s', got %s", expectedError, response["error"])
	}
}

func TestShortenURL_DuplicateURLs(t *testing.T) {
	req, err := http.NewRequest(
		"POST",
		"/shorten",
		bytes.NewBuffer([]byte(fmt.Sprintf(`{"originalUrl": "%s"}`, "https://localhost:8081"))),
	)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	w1 := httptest.NewRecorder()
	ShortenURL(w1, req)

	expectedStatusCode := http.StatusOK
	if w1.Code != expectedStatusCode {
		t.Errorf("expected status code %d, got %d", expectedStatusCode, w1.Code)
	}

	var response1 map[string]string
	err = json.Unmarshal(w1.Body.Bytes(), &response1)
	if err != nil {
		t.Fatal(err)
	}

	w2 := httptest.NewRecorder()
	ShortenURL(w2, req)

	if w2.Code != expectedStatusCode {
		t.Errorf("expected status code %d, got %d", expectedStatusCode, w2.Code)
	}

	var response2 map[string]string
	err = json.Unmarshal(w2.Body.Bytes(), &response2)
	if err != nil {
		t.Fatal(err)
	}

	if response1["shortUrl"] != response2["shortUrl"] {
		t.Errorf("expected the same alias for duplicate URLs, got %s and %s",
			response1["shortUrl"], response2["shortUrl"])
	}
}

func TestShortenURL_InvalidContentType(t *testing.T) {
	req, err := http.NewRequest(
		"POST",
		"/shorten",
		bytes.NewBuffer([]byte("not-a-json-body")),
	)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "text/plain")

	w := httptest.NewRecorder()
	ShortenURL(w, req)

	expectedStatusCode := http.StatusUnsupportedMediaType
	if w.Code != expectedStatusCode {
		t.Errorf("expected status code %d, got %d", expectedStatusCode, w.Code)
	}

	var response map[string]string
	err = json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	expectedError := "unsupported content type"
	if response["error"] != expectedError {
		t.Errorf("expected error message '%s', got %s", expectedError, response["error"])
	}
}
