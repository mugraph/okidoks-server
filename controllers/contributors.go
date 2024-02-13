package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mugraph/okidoks-server/models"
	"github.com/mugraph/okidoks-server/models/commonmeta"
)

// GET /contributors
func Contributors(c *gin.Context) {
	contributors := []commonmeta.Contributor{}

	models.DB.Preload("ContributorRoles").Preload("Resources").Find(&contributors)

	c.JSON(http.StatusOK, gin.H{"contributors": contributors})
}
