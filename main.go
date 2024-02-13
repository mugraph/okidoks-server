package main

import (
	"flag"

	"github.com/gin-gonic/gin"
	"github.com/mugraph/okidoks-server/controllers"
	"github.com/mugraph/okidoks-server/logger"
	"github.com/mugraph/okidoks-server/models"
	"github.com/mugraph/okidoks-server/models/datacite"
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
		log.Warn("could not get registration agency from DOI",
			"input", *doi,
			"error", err)
		return
	}

	if ra == "DataCite" {

		// Get DataCite attributes
		r, err := datacite.GetDataCite(*doi)
		if err != nil {
			log.Warn("could not get DataCite metadat for DOI",
				"input", *doi,
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

		// Add resource to DB
		models.DB.Create(&resource)
	}

	router.GET("/resource/:prefix/*suffix", controllers.FirstResource)
	router.GET("/resources", controllers.Resources)
	router.POST("/resources", controllers.CreateResource)

	router.GET("/contributors", controllers.Contributors)

	router.GET("/publishers", controllers.Publishers)

	router.Run(":8081")
}
