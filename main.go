package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mugraph/okidoks/okidoks-server/models"
)

func main() {
	fmt.Println("Hello World")

	router := gin.New()
	router.Use(CORSMiddleware())

	models.ConnectDatabase()

	router.GET("/publications", controllers.FindPublications)

	router.Run()
}
