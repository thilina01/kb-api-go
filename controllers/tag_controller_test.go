package controllers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/thilina01/kb-api-go/config"
)

func TestCreateTag(t *testing.T) {
	reqBody := []byte(`{"name": "test-tag"}`)
	req := httptest.NewRequest(http.MethodPost, "/tags", bytes.NewReader(reqBody))
	w := httptest.NewRecorder()

	err := config.ConnectDB()
	if err != nil {
		t.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	CreateTag(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

func TestListTags(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/tags", nil)
	w := httptest.NewRecorder()

	err := config.ConnectDB()
	if err != nil {
		t.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	ListTags(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}
