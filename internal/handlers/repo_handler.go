package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/Mariola04/Scalabit/internal/services"
)

func CreateRepo(c *gin.Context) {
	var req struct {
		Name string `json:"name" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil || req.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "required repo name"})
		return
	}

	client, ctx, err := services.NewGitHubClient()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error creating GitHub Client", "details": err.Error()})
		return
	}

	err = services.CreateRepository(client, ctx, req.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error creating repo", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "successfully created repo"})
}

func DeleteRepo(c *gin.Context) {
	owner := c.Param("owner")
	repo := c.Param("repo")

	if owner == "" || repo == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "owner and repo are mandatory"})
		return
	}

	client, ctx, err := services.NewGitHubClient()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error creating GitHub Client", "details": err.Error()})
		return
	}

	err = services.DeleteRepository(client, ctx, owner, repo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error deleting repo", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "successfully deleted repo"})
}

func ListRepos(c *gin.Context) {
	client, ctx, err := services.NewGitHubClient()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error creating GitHub Client", "details": err.Error()})
		return
	}

	repos, err := services.ListRepositories(client, ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error listing repos", "details": err.Error()})
		return
	}

	var result []gin.H
	for _, r := range repos {
		result = append(result, gin.H{
			"name": r.GetName(),
			"url":  r.GetHTMLURL(),
		})
	}

	c.JSON(http.StatusOK, result)
}
