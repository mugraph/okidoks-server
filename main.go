package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mugraph/okidoks-server/controllers"
	"github.com/mugraph/okidoks-server/models"
)

func main() {
	router := gin.New()

	models.ConnectDatabase()

	router.GET("/publications", controllers.FindPublications)

	router.Run(":8081")
}
