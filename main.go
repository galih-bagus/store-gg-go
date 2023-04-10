package main

import (
	"backend-storegg-go/controllers"
	"backend-storegg-go/helpers"

	"github.com/gin-gonic/gin"
)

func init() {
	helpers.LoadEnv()
	helpers.ConnectDb()
}

func main() {
	r := gin.Default()

	// Category
	r.GET("/category", controllers.CategoryIndex)
	r.POST("/category", controllers.CategoryStore)
	r.GET("/category/:id", controllers.CategoryShow)
	r.PUT("/category/:id", controllers.CategoryUpdate)
	r.DELETE("/category/:id", controllers.CategoryDestroy)

	// Nominal
	r.GET("/nominal", controllers.NominalIndex)
	r.POST("/nominal", controllers.NominalStore)
	r.GET("/nominal/:id", controllers.NominalShow)
	r.PUT("/nominal/:id", controllers.NominalUpdate)
	r.DELETE("/nominal/:id", controllers.NominalDestroy)

	r.Run()
}
