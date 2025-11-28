package main

import (
	"log"

	"github.com/badzboss/go-elasticsearch/controllers"
	"github.com/badzboss/go-elasticsearch/models"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.Use(gin.Logger())

	r.LoadHTMLGlob("templates/**/*")

	models.ConnectDatabase()
	models.DBMigrate()

	models.ESClientConnection()
	models.ESCreateIndexIfNotExists()

	r.GET("/blogs", controllers.BlogsIndex)
	r.GET("/blogs/:id", controllers.BlogsShow)

	r.POST("/blogs/index", controllers.BlogsBuildSerachIndex)

	log.Println("Server started")
	r.Run()
}
