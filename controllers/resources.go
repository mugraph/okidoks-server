package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mugraph/okidoks-server/models"
)

// GET /resources

func FindResources(c *gin.Context) {
	resources := []models.Resource{}

	models.DB.Find(&resources)

	c.JSON(http.StatusOK, &resources)
}
