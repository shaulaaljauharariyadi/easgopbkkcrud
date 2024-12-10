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
	// Kirim data ke template HTML
	c.HTML(http.StatusOK, "categorypage.html", gin.H{
		"category": categories,
	})
}

func AddCategory(c *gin.Context) {
	if c.Request.Method == "GET" {
		// Render the "addcategory" page
		c.HTML(200, "addcategory.html", gin.H{
			"title": "Add Category",
		})
	} else if c.Request.Method == "POST" {
		// Handle the form submission
		name := c.PostForm("name")
		// currentTime := time.Now()

		categories := models.Category{
			Name: string(name),
			// CreatedAt: currentTime,
			// UpdatedAt: currentTime,
		}

		result := initializers.DB.Create(&categories)

		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}

		c.Redirect(302, "/categorypage")
	}
}

// func Edit(c *gin.Context) {
// 	if c.Request.Method == http.MethodGet {
// 		id, err := strconv.Atoi(c.Query("id"))
// 		if err != nil {
// 			c.AbortWithStatus(http.StatusBadRequest)
// 			return
// 		}

// 		category := categorymodel.Detail(id)
// 		c.HTML(http.StatusOK, "category/edit.html", gin.H{
// 			"category": category,
// 		})
// 		return
// 	}

// 	id, _ := strconv.Atoi(c.PostForm("id"))
// 	var category entities.Category
// 	category.Name = c.PostForm("name")
// 	category.UpdatedAt = time.Now()

// 	if !categorymodel.Update(id, category) {
// 		c.Redirect(http.StatusSeeOther, c.Request.Referer())
// 		return
// 	}

// 	c.Redirect(http.StatusSeeOther, "/categories")
// }

// func Delete(c *gin.Context) {
// 	id, err := strconv.Atoi(c.Query("id"))
// 	if err != nil {
// 		c.AbortWithStatus(http.StatusBadRequest)
// 		return
// 	}

// 	if err := categorymodel.Delete(id); err != nil {
// 		c.AbortWithStatus(http.StatusInternalServerError)
// 		return
// 	}

// 	c.Redirect(http.StatusSeeOther, "/categories")
// }
