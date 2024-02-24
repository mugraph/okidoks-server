package main

import (
	"encoding/json"
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mugraph/okidoks-server/controllers"
	"github.com/mugraph/okidoks-server/logger"
	"github.com/mugraph/okidoks-server/models"
	"github.com/mugraph/okidoks-server/models/datacite"
	"github.com/mugraph/okidoks-server/utils"
)

var log = logger.Log

// GinModeToggle takes a pointer to a bool and sets gin debug mode on or off
func GinModeToggle(b *bool) {
	if b != nil && !*b {
		gin.SetMode(gin.ReleaseMode)
		log.Warn("set gin mode RELEASE mode", "input", *b)
		return
	}
	if b != nil && *b {
		gin.SetMode(gin.DebugMode)
		log.Warn("set gin to DEBUG mode", "input", *b)
		return
	}
}

func main() {
	// Parse command line flags
	doi := flag.String("doi", "", "a string")
	ginDebug := flag.Bool("ginDebug", true, "a bool")
	flag.Parse()

	GinModeToggle(ginDebug)

	router := gin.New()
	models.ConnectDatabase()

	if doi != nil && *doi != "" {
		ra, err := utils.GetDOIRA(*doi)
		if err != nil {
			log.Warn("could not get registration agency from DOI",
				"input", *doi,
				"error", err)
			return
		}

		if ra != "DataCite" {
			log.Warn(fmt.Sprintf("cannot read %s doi yet", ra), "input", *doi)
			return
		}

		if ra == "DataCite" || ra == "TestDataCite" {
			var r datacite.Resource
			if ra == "TestDataCite" {
				// Get DataCite attributes
				r, err = datacite.GetDataCite(*doi, true)
				if err != nil {
					log.Warn("could not get DataCite metadata for DOI",
						"input", *doi,
						"error", err)
					return
				}
			} else {
				// Get DataCite attributes
				r, err = datacite.GetDataCite(*doi, false)
				if err != nil {
					log.Warn("could not get DataCite metadata for DOI",
						"input", *doi,
						"error", err)
					return
				}
			}

			// Read DataCite into resource
			resource, err := datacite.ReadDataCite(r)
			if err != nil {
				log.Warn("could not read DataCite metadata to Resource",
					"error", err)
				return
			}

			// Print to Stdout
			res, _ := json.Marshal(resource.ToJSONModel())
			fmt.Println(string(res))
		}
	}

	router.GET("/api/v1/resource/:prefix/*suffix", controllers.FirstResource)
	router.GET("/api/v1/resources", controllers.Resources)
	router.GET("/api/v1/affiliations", controllers.Affiliations)
	router.POST("/api/v1/resources", controllers.CreateResource)

	router.GET("/api/v1/contributors", controllers.Contributors)

	router.GET("/api/v1/publishers", controllers.Publishers)

	router.Run(":8081")
}
