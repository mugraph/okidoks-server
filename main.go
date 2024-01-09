package main

import (
	"flag"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mugraph/okidoks-server/controllers"
	"github.com/mugraph/okidoks-server/models"
)

var debugMode bool

func main() {
	// Parse command line flags
	doi := flag.String("doi", "10.5555/12345678", "a string")
	debug := flag.Bool("debug", false, "Enable debug mode")
	flag.Parse()
	debugMode = *debug

	router := gin.New()

	models.ConnectDatabase()

	// Add Resources
	resource, err := models.ResourceFromDOI(*doi)
	if err != nil {
		log.Fatal(err)
	}
	models.AddResourceToDB(resource)

	router.GET("/resources", controllers.FindResources)
	router.GET("/authors", controllers.FindAuthors)
	router.GET("/affiliations", controllers.FindAffiliations)

	router.Run(":8081")
}
