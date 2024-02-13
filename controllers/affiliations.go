package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mugraph/okidoks-server/models"
	"github.com/mugraph/okidoks-server/models/commonmeta"
)

// GET /affiliations

func FindAffiliations(c *gin.Context) {
	affiliations := []commonmeta.Affiliation{}

	models.DB.Find(&affiliations)

	c.JSON(http.StatusOK, &affiliations)
}
