package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/Mariola04/Scalabit/internal/services"
)

func ListPullRequests(c *gin.Context) {
	owner := c.Param("owner")
	repo := c.Param("repo")
	nStr := c.DefaultQuery("n", "5")

	n, err := strconv.Atoi(nStr)
	if err != nil || n <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid 'n'"})
		return
	}

	client, ctx, err := services.NewGitHubClient()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error creating GitHub client", "details": err.Error()})
		return
	}

	prs, err := services.ListOpenPullRequests(client, ctx, owner, repo, n)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error listing public repos", "details": err.Error()})
		return
	}

	var result []gin.H
	for _, pr := range prs {
		result = append(result, gin.H{
			"title": pr.GetTitle(),
			"url":   pr.GetHTMLURL(),
			"user":  pr.GetUser().GetLogin(),
		})
	}

	c.JSON(http.StatusOK, result)
}
