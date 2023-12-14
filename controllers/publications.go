package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mugraph/okidoks-server/models"
)

// GET /publications

func FindPublications(c *gin.Context) {
	publications := []models.Publication{}

	models.DB.Find(&publications)

	c.JSON(http.StatusOK, &publications)
}
