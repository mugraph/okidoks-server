package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgconn"
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

type APIErrorResponse struct {
	Status  int    `json:"status,omitempty"`
	Success bool   `json:"success,omitempty"`
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
	Level   string `json:"level,omitempty"` // Info, Warn, Error
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

	err = models.DB.Where("id = ?", doiURL).
		Preload("Contributors.ContributorRoles").
		Preload(clause.Associations).
		First(&resource).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusOK, APIErrorResponse{
			Status:  http.StatusOK,
			Success: false,
			Message: "The resource you are looking for does not exist.",
			Level:   "Error",
		})
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
		c.JSON(http.StatusBadRequest, APIErrorResponse{
			Status:  http.StatusBadRequest,
			Success: false,
			Message: err.Error(),
			Level:   "Error",
		})
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
		r, err := datacite.GetDataCite(url, false)
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

		// Find or create license and assign its UUID to resource.License.UUID
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
			log.Warn("publisher record not found", "error", err)
		}

		resource.Publisher.UUID = publisher.UUID

		// Add resource to DB
		err = models.DB.Create(&resource).Error
		if err != nil {
			var pgErr *pgconn.PgError
			if errors.As(err, &pgErr) {
				log.Warn(pgErr.Message, "error", pgErr.Code)

				c.JSON(http.StatusOK, APIErrorResponse{
					Status:  http.StatusOK,
					Success: false,
					Message: utils.WithFullStop(utils.FirstToUpper(pgErr.Message)),
					Error:   "SQLSTATE " + pgErr.Code,
					Level:   "Warn",
				})

				return
			}
		}

		c.JSON(http.StatusOK, gin.H{"data": &resource})
	}
}
