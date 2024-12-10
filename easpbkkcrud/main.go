package main

import (
	"easpbkkcrudnew/easpbkkcrud/controllers"
	"easpbkkcrudnew/easpbkkcrud/initializers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	initializers.ConnectDatabase()

	router := gin.Default()

	router.LoadHTMLGlob("views/*")

	router.GET("/homepage", controllers.Welcome)

	router.GET("/categorypage", controllers.CategoryPage) // Index
	router.GET("/addcategory", controllers.AddCategory)   // Add (GET)
	router.POST("/addcategory", controllers.AddCategory)  // Add (POST)

	router.GET("/productpage", controllers.ProductPage) // Index
	router.GET("/addproduct", controllers.AddProduct)   // Add (GET)
	router.POST("/addproduct", controllers.AddProduct)  // Add (POST)

	log.Println("Server running on port 8080")
	router.Run(":8080")
}
