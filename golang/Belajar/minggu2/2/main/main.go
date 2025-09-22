package main

import (
	"2/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Product struct {
Name string `json:"name" binding:"required,min=3"`
Price float64 `json:"price" binding:"required,gt=0"`
Stock int `json:"stock" binding:"required,gte=0"`
}

func main() {
	r := gin.Default()
	r.Use(middleware.ErrorHandler)

	r.POST("/products", func(c *gin.Context) {
		var body Product
		if err := c.ShouldBindJSON(&body); err != nil {
			c.Error(err)
			return
		}
		c.JSON(http.StatusCreated, gin.H{"message": "produk valid", "data": body})
	})

	r.Run(":8080")
}