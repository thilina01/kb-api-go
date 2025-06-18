package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/thilina01/kb-api-go/config"
)

func TestGetArticleByIDInvalid(t *testing.T) {
	err := config.ConnectDB()
	if err != nil {
		t.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/articles/invalid-id", nil)
	w := httptest.NewRecorder()

	GetArticleByID(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400, got %d", w.Code)
	}
}

func TestDeleteArticleInvalidID(t *testing.T) {
	err := config.ConnectDB()
	if err != nil {
		t.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	req := httptest.NewRequest(http.MethodDelete, "/articles/invalid-id", nil)
	w := httptest.NewRecorder()

	DeleteArticle(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400, got %d", w.Code)
	}
}

func TestSearchArticlesNoQuery(t *testing.T) {
	err := config.ConnectDB()
	if err != nil {
		t.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/articles/search", nil)
	w := httptest.NewRecorder()

	SearchArticles(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400, got %d", w.Code)
	}
}
