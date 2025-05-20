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
		c.JSON(http.StatusBadRequest, gin.H{"error": "parâmetro 'n' inválido"})
		return
	}

	client, ctx, err := services.NewGitHubClient()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao criar cliente GitHub", "details": err.Error()})
		return
	}

	prs, err := services.ListOpenPullRequests(client, ctx, owner, repo, n)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao listar PRs", "details": err.Error()})
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
