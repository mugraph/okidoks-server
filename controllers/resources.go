package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mugraph/okidoks-server/logger"
	"github.com/mugraph/okidoks-server/models"
	"github.com/mugraph/okidoks-server/utils"
)

var log = logger.Log

// GET /resources
func FindResources(c *gin.Context) {
	resources := []models.Resource{}

	models.DB.Preload("ContributorRoles").Preload("Contributors").Find(&resources)

	c.JSON(http.StatusOK, &resources)
}

// POST /resources
// Create new Resource
func CreateResource(c *gin.Context) {

	//Validate Input
	var dp doiPayload
	if err := c.ShouldBindJSON(&dp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ra, err := utils.GetDOIRA(dp.URL)
	if err != nil {
		log.Warn("Could not get registration agency from DOI: %v. Error:", dp.URL, err)
	} else if ra == "DataCite" {

		// Get DataCite attributes
		attr, err := models.GetDataCite(dp.URL)
		if err != nil {
			log.Warn("Could not get DataCite metadat for DOI: %v. Error:", dp.URL, err)
		}

		// Read DataCite into resource
		resource, err := models.ReadDataCite(attr)
		if err != nil {
			log.Warn("Could not read DataCite metadata to Resource. Error:", err)
		}

		// Add resource to DB
		models.DB.Create(&resource)

		c.JSON(http.StatusOK, gin.H{"data": resource})
	}
}
