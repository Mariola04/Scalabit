
package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/Mariola04/Scalabit/internal/handlers"
)

// helper function to set up a router with routes
func setupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/repos", handlers.CreateRepo)
	r.DELETE("/repos/:owner/:repo", handlers.DeleteRepo)
	r.GET("/repos", handlers.ListRepos)
	r.GET("/repos/:owner/:repo/pulls", handlers.ListPullRequests)
	return r
}

func TestCreateRepo_Valid(t *testing.T) {
	router := setupRouter()

	body := map[string]string{"name": "test-repo"}
	jsonBody, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", "/repos", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated && w.Code != http.StatusInternalServerError {
		t.Errorf("expected status 201 or 500, got %d", w.Code)
	}
}

func TestCreateRepo_MissingName(t *testing.T) {
	router := setupRouter()

	body := map[string]string{}
	jsonBody, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", "/repos", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected status 400, got %d", w.Code)
	}
}

func TestDeleteRepo_MissingParams(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("DELETE", "/repos//missing", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest && w.Code != http.StatusNotFound && w.Code != http.StatusInternalServerError {
	t.Errorf("expected status 400, 404 or 500, got %d", w.Code)
	}
}

func TestListRepos(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("GET", "/repos", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK && w.Code != http.StatusInternalServerError {
		t.Errorf("expected status 200 or 500, got %d", w.Code)
	}
}

func TestListPullRequests_InvalidN(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("GET", "/repos/test/test/pulls?n=abc", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected status 400 for invalid 'n', got %d", w.Code)
	}
}