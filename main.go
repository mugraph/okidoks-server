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

	// Get DataCite attributes
	attr, err := models.GetDataCite(*doi)
	if err != nil {
		log.Fatal(err)
	}

	// Read DataCite into resource
	resource, err := models.ReadDataCite(attr)
	if err != nil {
		log.Fatal(err)
	}
	models.DB.Create(&resource)

	router.GET("/resources", controllers.FindResources)

	router.Run(":8081")
}
