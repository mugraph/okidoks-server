package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mugraph/okidoks-server/logger"
	"github.com/mugraph/okidoks-server/models"
	"github.com/mugraph/okidoks-server/utils"
	"gorm.io/gorm/clause"
)

var log = logger.Log

type doiPayload struct {
	URL string
}

type resourcePayload struct {
	Resources []models.Resource `json:"resources"`
}

// GET /resources
func FindResources(c *gin.Context) {
	resources := []models.Resource{}

	models.DB.Preload("Contributors.ContributorRoles").Preload(clause.Associations).Find(&resources)

	c.JSON(http.StatusOK, resourcePayload{
		Resources: resources,
	})
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

	doiURL, err := utils.DOIAsURL(dp.URL)
	if err != nil {
		log.Warn("could not get DOI as URL", "input", dp.URL, "error", err)
		return
	}

	ra, err := utils.GetDOIRA(doiURL)
	if err != nil {
		log.Warn("could not get registration agency from DOI", "input", doiURL, "error", err)
		return
	}

	if ra == "DataCite" {

		// Get DataCite attributes
		attr, err := models.GetDataCite(doiURL)
		if err != nil {
			log.Warn("could not get DataCite metadata for DOI", "input", doiURL, "error", err)
			return
		}
		// Read DataCite into resource
		resource, err := models.ReadDataCite(attr)
		if err != nil {
			log.Warn("could not read DataCite metadata to Resource", "error", err)
			return
		}

		// Add resource to DB
		models.DB.Create(&resource)

		c.JSON(http.StatusOK, gin.H{"data": resource})

	}
}
