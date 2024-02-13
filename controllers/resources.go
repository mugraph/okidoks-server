package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mugraph/okidoks-server/logger"
	"github.com/mugraph/okidoks-server/models"
	"github.com/mugraph/okidoks-server/models/commonmeta"
	"github.com/mugraph/okidoks-server/models/datacite"
	"github.com/mugraph/okidoks-server/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var log = logger.Log

type doiPayload struct {
	URL string
}

type StatusError struct {
	Status int    `json:"status"`
	Title  string `json:"title"`
}

// GET /resource/:prefix/*suffix
func FirstResource(c *gin.Context) {
	var resource commonmeta.Resource

	prefix := c.Param("prefix")
	suffix := c.Param("suffix")

	doiURL, err := utils.DOIAsURL(prefix + suffix)
	if err != nil {
		log.Warn("could not get DOI as URL")
	}

	res := models.DB.Where("id = ?", doiURL).
		Preload("Contributors.ContributorRoles").
		Preload(clause.Associations).
		First(&resource)

	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		err := []StatusError{}
		err = append(err, StatusError{
			Status: http.StatusNotFound,
			Title:  "The resource you are looking for doesn't exist.",
		})
		c.JSON(http.StatusNotFound, gin.H{"errors": err})
		return
	}

	// Pass address as custom MarshalJSON method is defined on pointer receiver.
	c.JSON(http.StatusOK, gin.H{"resource": &resource})
}

// GET /resources
func Resources(c *gin.Context) {
	var resources []commonmeta.Resource

	models.DB.Preload("Contributors.ContributorRoles").
		Preload(clause.Associations).
		Find(&resources)

	// json.Marshal traverses values recursively and calls custom MarshalJSON
	// method where possible.
	c.JSON(http.StatusOK, gin.H{"resources": resources})
}

// FindCreateLicense takes a pointer to a commonmeta.License struct and queries
// the DB License table for the first entry that matches the URL of the provided
// License struct. It creates the License in the DB in case no matching record was
// found and calls itself for another attempt.
// It returns the retreived commonmeta.License if the query was successful.
func FindOrCreateLicense(rl *commonmeta.License) commonmeta.License {
	var l commonmeta.License
	err := models.DB.Where("URL = ?", rl.URL).First(&l).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		log.Warn("license record not found", "error", err)
		
		models.DB.Create(&rl)
		
		l = FindOrCreateLicense(rl)
	}
	return l
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

	url, err := utils.DOIAsURL(dp.URL)
	if err != nil {
		log.Warn("could not get DOI as URL", "input", dp.URL, "error", err)
		return
	}

	ra, err := utils.GetDOIRA(url)
	if err != nil {
		log.Warn("could not get registration agency from DOI",
			"input", url,
			"error", err)
		return
	}

	if ra == "DataCite" {
		// Get DataCite attributes
		r, err := datacite.GetDataCite(url)
		if err != nil {
			log.Warn("could not get DataCite metadata for DOI",
				"input", url,
				"error", err)
			return
		}

		// Read DataCite into resource
		resource, err := datacite.ReadDataCite(r)
		if err != nil {
			log.Warn("could not read DataCite metadata to Resource",
				"error", err)
			return
		}

		// Find or create License and assign its UUID to resource.License.UUID
		// to comply with foreignKey contraints
		if resource.License != nil {
			license := FindOrCreateLicense(resource.License)
			resource.License.UUID = license.UUID
		}

		// Find Publisher by its unique index and assign its UUID to resource.License.UUID
		// to comply with foreignKey contraints 
		var publisher commonmeta.Publisher	 
		err = models.DB.Where("Name = ?", resource.Publisher.Name).First(&publisher).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Warn("license record not found", "error", err)
		}
		resource.Publisher.UUID = publisher.UUID

		// Add resource to DB
		models.DB.Create(&resource)

		c.JSON(http.StatusOK, gin.H{"data": resource})
	}
}
