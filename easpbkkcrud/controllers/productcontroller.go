package controllers

import (
	"easpbkkcrudnew/easpbkkcrud/initializers"
	"easpbkkcrudnew/easpbkkcrud/models"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ProductPage(c *gin.Context) {
	var products []models.Product
	product := initializers.DB.Preload("Category").Find(&products)

	if product.Error != nil {
		log.Println("Error fetching products:", product.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
		return
	}

	log.Println("Products:", products)

	c.HTML(http.StatusOK, "productpage.html", gin.H{
		"product": products,
	})
}

func DetailProduct(c *gin.Context) {
	id := c.Param("id")
	productID, err := strconv.Atoi(id)
	if err != nil {
		log.Println("Invalid product ID:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	var products models.Product
	result := initializers.DB.Preload("Category").First(&products, productID)

	if result.Error != nil {
		log.Println("Error fetching product details:", result.Error)
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	log.Println("Product Details:", products)

	c.HTML(http.StatusOK, "detailproduct.html", gin.H{
		"product": products,
	})
}

func AddProduct(c *gin.Context) {
	if c.Request.Method == "GET" {
		var categories []models.Category
		if err := initializers.DB.Find(&categories).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch categories"})
			return
		}

		c.HTML(http.StatusOK, "addproduct.html", gin.H{
			"products": categories,
		})

		return
	}

	if c.Request.Method == "POST" {
		name := c.PostForm("name")
		category_id, _ := strconv.ParseUint(c.PostForm("category_id"), 10, 64)
		fmt.Println(category_id)
		stock, _ := strconv.ParseFloat(c.PostForm("stock"), 64)
		description := c.PostForm("description")

		products := models.Product{
			Name:        string(name),
			Category_Id: uint(category_id),
			Stock:       int(stock),
			Description: string(description),
		}
		result := initializers.DB.Create(&products)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}

		c.Redirect(http.StatusFound, "/productpage")
	}
}

func EditProduct(c *gin.Context) {
	if c.Request.Method == "GET" {
		id := c.Param("id")
		productID, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
			return
		}

		var product models.Product
		if err := initializers.DB.First(&product, productID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}

		var categories []models.Category
		if err := initializers.DB.Find(&categories).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch categories"})
			return
		}

		c.HTML(http.StatusOK, "editproduct.html", gin.H{
			"product":    product,
			"categories": categories,
		})
	}

	if c.Request.Method == "POST" {
		id := c.Param("id")
		productID, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
			return
		}

		name := c.PostForm("name")
		categoryID, _ := strconv.Atoi(c.PostForm("category_id"))
		stock, _ := strconv.Atoi(c.PostForm("stock"))
		description := c.PostForm("description")

		var product models.Product
		if err := initializers.DB.First(&product, productID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}

		product.Name = name
		product.Category_Id = uint(categoryID)
		product.Stock = stock
		product.Description = description

		if err := initializers.DB.Save(&product).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
			return
		}

		c.Redirect(http.StatusFound, "/productpage")
	}
}

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	productID, err := strconv.Atoi(id)
	if err != nil {
		log.Println("Invalid product ID:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	var product models.Product
	if err := initializers.DB.First(&product, productID).Error; err != nil {
		log.Println("Product not found:", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	if err := initializers.DB.Delete(&product).Error; err != nil {
		log.Println("Error deleting product:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
		return
	}

	log.Println("Product deleted successfully:", product)
	c.Redirect(http.StatusFound, "/productpage")
}
