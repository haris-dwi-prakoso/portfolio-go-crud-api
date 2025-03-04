package main

import (
	"harisdwiprakoso/crud-api/controllers"
	"harisdwiprakoso/crud-api/models"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	models.ConnectDatabase()

	router.POST("/products", controllers.CreateProduct)
	router.GET("/products", controllers.FindProducts)
	router.GET("/products/:id", controllers.FindProduct)
	router.PATCH("/products/:id", controllers.UpdateProduct)
	router.DELETE("/products/:id", controllers.DeleteProduct)

	router.Run("localhost:8080")
}
