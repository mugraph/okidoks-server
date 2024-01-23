package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mugraph/okidoks-server/models"
)

// GET /contributors
func FindContributors(c *gin.Context) {
	contributors := []models.Contributor{}

	models.DB.Preload("ContributorRoles").Preload("Resources").Find(&contributors)

	c.JSON(http.StatusOK, &contributors)
}

type doiPayload struct {
	URL string
}
