package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/Mariola04/Scalabit/internal/services"
)


/*
In Gin, every HTTP request is handled by a function called a handler, and Gin gives that function a *gin.Context object as a parameter.

This object (*gin.Context, usually called c) gives you full access to:

    The request (method, URL, params, headers, body, etc...)

    The response (status code, JSON response, etc...)

    Middleware and other request-scoped values

    *gin.Context is like the "interface" between my code and the incoming HTTP request/response.
*/

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
