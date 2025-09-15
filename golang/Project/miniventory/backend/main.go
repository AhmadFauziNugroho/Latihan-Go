package main

import (
	"log"

	"miniventory/backend/database"
	"miniventory/backend/handlers"
	"miniventory/backend/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()
	if err := database.DB.AutoMigrate(&models.Product{}); err != nil {
		log.Fatalf("gagal migrasi: %v", err)
	}

	r := gin.Default()
	r.Use(cors.Default())

	api := r.Group("/api")
	{
		api.GET("/products", handlers.GetAllProducts)
		api.POST("/products", handlers.CreateProduct)
		api.PUT("/products/:id", handlers.UpdateProduct)
		api.DELETE("/products/:id", handlers.DeleteProduct)
	}

	r.Run(":8080")
}