package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mugraph/okidoks-server/models"
)

// GET /publishers
func Publishers(c *gin.Context) {
	publishers := []models.Publisher{}

	models.DB.Find(&publishers)

	c.JSON(http.StatusOK, gin.H{"publishers": publishers})
}
