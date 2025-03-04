package controllers

import (
	"harisdwiprakoso/crud-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateProductInput struct {
	Name        string  `json:"name" binding:"required"`
	Description *string `json:"description"`
	Price       int     `json:"price" binding:"required"`
	Stock       int     `json:"stock" binding:"required"`
}

type UpdateProductInput struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
	Price       int     `json:"price"`
	Stock       int     `json:"stock"`
}

func CreateProduct(c *gin.Context) {
	var input CreateProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product := models.Product{Name: input.Name, Description: input.Description, Price: input.Price, Stock: input.Stock}
	models.DB.Create(&product)

	c.JSON(http.StatusOK, gin.H{"data": product})
}

func FindProducts(c *gin.Context) {
	var products []models.Product
	models.DB.Find(&products)

	c.JSON(http.StatusOK, gin.H{"data": products})
}

func FindProduct(c *gin.Context) {
	var product models.Product

	if err := models.DB.Where("id = ?", c.Param("id")).First((&product)).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": product})
}

func UpdateProduct(c *gin.Context) {
	var product models.Product

	if err := models.DB.Where("id = ?", c.Param("id")).First((&product)).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	var input UpdateProductInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedProduct := models.Product{Name: input.Name, Description: input.Description, Price: input.Price, Stock: input.Stock}

	models.DB.Model(&product).Updates(&updatedProduct)
	c.JSON(http.StatusOK, gin.H{"data": product})
}

func DeleteProduct(c *gin.Context) {
	var product models.Product

	if err := models.DB.Where("id = ?", c.Param("id")).First((&product)).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	models.DB.Delete(&product)
	c.JSON(http.StatusOK, gin.H{"data": "Successfully deleted"})
}
