package controllers

import (
	"easpbkkcrudnew/easpbkkcrud/initializers"
	"easpbkkcrudnew/easpbkkcrud/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ProductPage(c *gin.Context) {
	var products []models.Product
	items := initializers.DB.Find(&products)

	if items.Error != nil {
		log.Println("Error fetching products:", items.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
		return
	}

	log.Println("Products:", products)
	// Kirim data ke template HTML
	c.HTML(http.StatusOK, "productpage.html", gin.H{
		"product": products,
	})
}

// func Detail(c *gin.Context) {
// 	id, err := strconv.Atoi(c.Query("id"))
// 	if err != nil {
// 		c.AbortWithStatus(http.StatusBadRequest)
// 		return
// 	}

// 	product := mode.Detail(id)
// 	c.HTML(http.StatusOK, "product/detail.html", gin.H{
// 		"product": product,
// 	})
// }

func AddProduct(c *gin.Context) {
	if c.Request.Method == "GET" {
		var categories []models.Category
		if err := initializers.DB.Find(&categories).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch categories"})
			return
		}

		c.HTML(http.StatusOK, "addproduct.html", gin.H{
			"category": categories,
		})

		return
	}

	if c.Request.Method == "POST" {
		name := c.PostForm("name")
		category_id, _ := strconv.ParseUint(c.PostForm("category_id"), 10, 64)
		// fmt.Println(category_id)
		stock, _ := strconv.ParseFloat(c.PostForm("stock"), 64)
		description := c.PostForm("description")
		// currentTime := time.Now()

		products := models.Product{
			Name:        string(name),
			Category_Id: uint(category_id),
			Stock:       int(stock),
			Description: string(description),
			// CreatedAt:   currentTime,
			// UpdatedAt:   currentTime,
		}
		result := initializers.DB.Create(&products)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}

		c.Redirect(http.StatusFound, "/productpage")
	}
}

// func Edit(c *gin.Context) {
// 	if c.Request.Method == http.MethodGet {
// 		id, err := strconv.Atoi(c.Query("id"))
// 		if err != nil {
// 			c.AbortWithStatus(http.StatusBadRequest)
// 			return
// 		}

// 		product := productmodel.Detail(id)
// 		categories := categorymodel.GetAll()
// 		c.HTML(http.StatusOK, "product/edit.html", gin.H{
// 			"product":    product,
// 			"categories": categories,
// 		})
// 		return
// 	}

// 	id, _ := strconv.Atoi(c.PostForm("id"))
// 	categoryId, _ := strconv.Atoi(c.PostForm("category_id"))
// 	stock, _ := strconv.Atoi(c.PostForm("stock"))

// 	var product entities.Product
// 	product.Name = c.PostForm("name")
// 	product.Category.Id = uint(categoryId)
// 	product.Stock = int64(stock)
// 	product.Description = c.PostForm("description")
// 	product.UpdatedAt = time.Now()

// 	if !productmodel.Update(id, product) {
// 		c.Redirect(http.StatusTemporaryRedirect, c.Request.Referer())
// 		return
// 	}

// 	c.Redirect(http.StatusSeeOther, "/products")
// }

// func Delete(c *gin.Context) {
// 	id, err := strconv.Atoi(c.Query("id"))
// 	if err != nil {
// 		c.AbortWithStatus(http.StatusBadRequest)
// 		return
// 	}

// 	if err := productmodel.Delete(id); err != nil {
// 		c.AbortWithStatus(http.StatusInternalServerError)
// 		return
// 	}

// 	c.Redirect(http.StatusSeeOther, "/products")
// }
