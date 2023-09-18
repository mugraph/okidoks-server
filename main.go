package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mugraph/okidoks-server/models"
	"github.com/mugraph/okidoks-server/controllers"
)

func main() {
	router := gin.New()

	models.ConnectDatabase()

	router.GET("/publications", controllers.FindPublications)

	router.Run()
}
