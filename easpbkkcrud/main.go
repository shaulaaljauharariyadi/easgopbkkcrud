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

	router.GET("/categorypage", controllers.CategoryPage)
	router.GET("/addcategory", controllers.AddCategory)
	router.POST("/addcategory", controllers.AddCategory)
	router.GET("/editcategory/:id", controllers.EditCategory)
	router.POST("/editcategory/:id", controllers.EditCategory)
	router.POST("/deletecategory/:id", controllers.DeleteCategory)

	router.GET("/productpage", controllers.ProductPage)
	router.GET("/addproduct", controllers.AddProduct)
	router.POST("/addproduct", controllers.AddProduct)
	router.GET("/detailproduct/:id", controllers.DetailProduct)
	router.GET("/editproduct/:id", controllers.EditProduct)
	router.POST("/editproduct/:id", controllers.EditProduct)
	router.POST("/deleteproduct/:id", controllers.DeleteProduct)

	log.Println("Server running on port 8080")
	router.Run(":8080")
}
