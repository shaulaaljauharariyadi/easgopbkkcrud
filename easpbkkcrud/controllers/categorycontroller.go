package controllers

import (
	"easpbkkcrudnew/easpbkkcrud/initializers"
	"easpbkkcrudnew/easpbkkcrud/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CategoryPage(c *gin.Context) {
	var categories []models.Category
	products := initializers.DB.Find(&categories)

	if products.Error != nil {
		log.Println("Error fetching products:", products.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
		return
	}

	log.Println("Categories:", categories)

	c.HTML(http.StatusOK, "categorypage.html", gin.H{
		"products": categories,
	})
}

func AddCategory(c *gin.Context) {
	if c.Request.Method == "GET" {
		c.HTML(200, "addcategory.html", gin.H{
			"title": "Add Category",
		})
	} else if c.Request.Method == "POST" {
		name := c.PostForm("name")

		categories := models.Category{
			Name: string(name),
		}

		result := initializers.DB.Create(&categories)

		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}

		c.Redirect(302, "/categorypage")
	}
}

func EditCategory(c *gin.Context) {
	if c.Request.Method == "GET" {
		id := c.Param("id")
		var category models.Category

		if err := initializers.DB.First(&category, id).Error; err != nil {
			log.Println("Error fetching category:", err)
			c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
			return
		}

		c.HTML(http.StatusOK, "editcategory.html", gin.H{
			"category": category,
		})
		return
	}

	if c.Request.Method == "POST" {
		id := c.PostForm("id")
		name := c.PostForm("name")

		if err := initializers.DB.Model(&models.Category{}).Where("id = ?", id).Update("name", name).Error; err != nil {
			log.Println("Error updating category:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update category"})
			return
		}

		c.Redirect(http.StatusFound, "/categorypage")
	}
}

func DeleteCategory(c *gin.Context) {
	id := c.Param("id")

	if err := initializers.DB.Delete(&models.Category{}, id).Error; err != nil {
		log.Println("Error deleting category:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete category"})
		return
	}

	c.Redirect(http.StatusFound, "/categorypage")
}
