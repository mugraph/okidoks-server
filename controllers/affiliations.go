package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mugraph/okidoks-server/models"
)

// GET /affiliations

func FindAffiliations(c *gin.Context) {
	affiliations := []models.Affiliation{}

	models.DB.Find(&affiliations)

	c.JSON(http.StatusOK, &affiliations)
}
