package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/thilina01/kb-api-go/config"
)

func TestCreateAndGetArticle(t *testing.T) {
	err := config.ConnectDB()
	if err != nil {
		t.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Create tag first
	tagReq := httptest.NewRequest(http.MethodPost, "/tags", strings.NewReader(`{"name":"test"}`))
	tagRes := httptest.NewRecorder()
	CreateTag(tagRes, tagReq)
	if tagRes.Code != http.StatusOK {
		t.Fatalf("Expected 200 creating tag, got %d", tagRes.Code)
	}
	var tagResp map[string]interface{}
	_ = json.Unmarshal(tagRes.Body.Bytes(), &tagResp)
	tagID := tagResp["id"].(string)

	// Create article
	articleBody := `{"title": "Test Article", "content": "Some content", "tags": ["` + tagID + `"]}`
	req := httptest.NewRequest(http.MethodPost, "/articles", strings.NewReader(articleBody))
	w := httptest.NewRecorder()
	CreateArticle(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("Expected 200 creating article, got %d", w.Code)
	}
	var articleResp map[string]interface{}
	_ = json.Unmarshal(w.Body.Bytes(), &articleResp)
	id := articleResp["id"].(string)

	// Fetch article by ID
	getReq := httptest.NewRequest(http.MethodGet, "/articles/"+id, nil)
	getW := httptest.NewRecorder()
	GetArticleByID(getW, getReq)
	if getW.Code != http.StatusOK {
		t.Errorf("Expected 200 fetching article, got %d", getW.Code)
	}
}

func TestUpdateArticle(t *testing.T) {
	err := config.ConnectDB()
	if err != nil {
		t.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Create article
	articleBody := `{"title": "ToUpdate", "content": "content", "tags": []}`
	req := httptest.NewRequest(http.MethodPost, "/articles", strings.NewReader(articleBody))
	w := httptest.NewRecorder()
	CreateArticle(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("Failed to create article")
	}
	var created map[string]interface{}
	_ = json.Unmarshal(w.Body.Bytes(), &created)
	id := created["id"].(string)

	// Update article
	updatedBody := `{"title": "Updated Title", "content": "Updated content", "tags": []}`
	putReq := httptest.NewRequest(http.MethodPut, "/articles/"+id, strings.NewReader(updatedBody))
	putW := httptest.NewRecorder()
	UpdateArticle(putW, putReq)
	if putW.Code != http.StatusOK {
		t.Errorf("Expected 200 updating article, got %d", putW.Code)
	}
}

func TestDeleteArticle(t *testing.T) {
	err := config.ConnectDB()
	if err != nil {
		t.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	articleBody := `{"title": "ToDelete", "content": "remove me", "tags": []}`
	req := httptest.NewRequest(http.MethodPost, "/articles", strings.NewReader(articleBody))
	w := httptest.NewRecorder()
	CreateArticle(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("Failed to create article")
	}
	var created map[string]interface{}
	_ = json.Unmarshal(w.Body.Bytes(), &created)
	id := created["id"].(string)

	// Delete it
	delReq := httptest.NewRequest(http.MethodDelete, "/articles/"+id, nil)
	delW := httptest.NewRecorder()
	DeleteArticle(delW, delReq)
	if delW.Code != http.StatusOK {
		t.Errorf("Expected 200 deleting article, got %d", delW.Code)
	}
}

func TestListArticlesAndSearch(t *testing.T) {
	err := config.ConnectDB()
	if err != nil {
		t.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// List articles
	listReq := httptest.NewRequest(http.MethodGet, "/articles", nil)
	listW := httptest.NewRecorder()
	ListArticles(listW, listReq)
	if listW.Code != http.StatusOK {
		t.Errorf("Expected 200 listing articles, got %d", listW.Code)
	}

	// Basic search
	searchReq := httptest.NewRequest(http.MethodGet, "/articles/search?q=test", nil)
	searchW := httptest.NewRecorder()
	SearchArticles(searchW, searchReq)
	if searchW.Code != http.StatusOK {
		t.Errorf("Expected 200 searching articles, got %d", searchW.Code)
	}
}
