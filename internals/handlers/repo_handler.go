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
		c.JSON(http.StatusBadRequest, gin.H{"error": "nome do repositório é obrigatório"})
		return
	}

	err := services.CreateRepository(req.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao criar repositório", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "repositório criado com sucesso"})
}

func DeleteRepo(c *gin.Context) {
	owner := c.Param("owner")
	repo := c.Param("repo")

	if owner == "" || repo == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "owner e repo são obrigatórios"})
		return
	}

	err := services.DeleteRepository(owner, repo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao apagar repositório", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "repositório removido com sucesso"})
}

func ListRepos(c *gin.Context) {
	repos, err := services.ListRepositories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao listar repositórios", "details": err.Error()})
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
