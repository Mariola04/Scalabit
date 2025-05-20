package tests

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/Mariola04/Scalabit/internal/handlers"
	"github.com/stretchr/testify/assert"
)


func TestCreateRepo_InvalidInput(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/repos", handlers.CreateRepo)

	body := []byte(`{}`) 
	req, _ := http.NewRequest(http.MethodPost, "/repos", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestDeleteRepo_MissingParams(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.DELETE("/repos/:owner/:repo", handlers.DeleteRepo)

	// Tests valid route but no paramethers
	req, _ := http.NewRequest(http.MethodDelete, "/repos//fake", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}


func TestListRepos_WithoutToken(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/repos", handlers.ListRepos)

	req, _ := http.NewRequest(http.MethodGet, "/repos", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
}

func TestListPullRequests_InvalidN(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/repos/:owner/:repo/pulls", handlers.ListPullRequests)

	req, _ := http.NewRequest(http.MethodGet, "/repos/test/test/pulls?n=abc", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}
