package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mugraph/okidoks-server/models"
)

// GET /authors
func FindAuthors(c *gin.Context) {
	authors := []models.Author{}

	models.DB.Find(&authors)

	c.JSON(http.StatusOK, &authors)
}
