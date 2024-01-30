package main

import (
	"flag"

	"github.com/gin-gonic/gin"
	"github.com/mugraph/okidoks-server/controllers"
	"github.com/mugraph/okidoks-server/logger"
	"github.com/mugraph/okidoks-server/models"
	"github.com/mugraph/okidoks-server/utils"
)

var log = logger.Log

func main() {
	log.Info("Hello, World!")

	// Parse command line flags
	doi := flag.String("doi", "10.5555/12345678", "a string")
	flag.Parse()

	router := gin.New()
	models.ConnectDatabase()

	ra, err := utils.GetDOIRA(*doi)
	if err != nil {
		log.Warn("could not get registration agency from DOI", "input", *doi, "error", err)
		return
	}

	if ra == "DataCite" {

		// Get DataCite attributes
		attr, err := models.GetDataCite(*doi)
		if err != nil {
			log.Warn("could not get DataCite metadat for DOI", "input", *doi, "error", err)
		}

		// Read DataCite into resource
		resource, err := models.ReadDataCite(attr)
		if err != nil {
			log.Warn("could not read DataCite metadata to Resource", "error", err)
		}

		// Add resource to DB
		models.DB.Create(&resource)
	}

	router.GET("/resources", controllers.FindResources)
	router.POST("/resources", controllers.CreateResource)

	router.GET("/contributors", controllers.FindContributors)

	router.GET("/publishers", controllers.FindPublishers)

	router.Run(":8081")
}
