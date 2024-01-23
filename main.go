package main

import (
	"flag"

	"github.com/gin-gonic/gin"
	"github.com/mugraph/okidoks-server/controllers"
	"github.com/mugraph/okidoks-server/logger"
	"github.com/mugraph/okidoks-server/models"
	"github.com/mugraph/okidoks-server/utils"
)

var debugMode bool
var log = logger.Log

func main() {
	log.Info("Hello, World!")

	// Parse command line flags
	doi := flag.String("doi", "10.5555/12345678", "a string")
	debug := flag.Bool("debug", false, "Enable debug mode")
	flag.Parse()
	debugMode = *debug

	router := gin.New()
	models.ConnectDatabase()

	ra, err := utils.GetDOIRA(*doi)
	if err != nil {
		log.Warn("Could not get registration agency from DOI: %v. Error:", *doi, err)
	} else if ra == "DataCite" {

		// Get DataCite attributes
		attr, err := models.GetDataCite(*doi)
		if err != nil {
			log.Warn("Could not get DataCite metadat for DOI: %v. Error:", *doi, err)
		}

		// Read DataCite into resource
		resource, err := models.ReadDataCite(attr)
		if err != nil {
			log.Warn("Could not read DataCite metadata to Resource. Error:", err)
		}

		// Add resource to DB
		models.DB.Create(&resource)
	}

	router.GET("/resources", controllers.FindResources)
	router.POST("/resources", controllers.CreateResource)

	router.GET("/contributors", controllers.FindContributors)

	router.Run(":8081")
}
