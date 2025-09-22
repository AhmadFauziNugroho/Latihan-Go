package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	ID uint `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Price float64 `json:"price"`
	Stock int `json:"stock"`
}

var DB *gorm.DB

func main() {
	db, _ := gorm.Open(sqlite.Open("app.db"), &gorm.Config{}) 
	db.AutoMigrate(&Product{})
	DB = db
	r := gin.Default()

	r.POST("/products", func(c *gin.Context) {
		var body Product
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		DB.Create(&body)
		c.JSON(http.StatusCreated, body)
	})


	r.GET("/products", func(c *gin.Context) {
		var products []Product
		DB.Find(&products)
		c.JSON(http.StatusOK, products)
	})

	r.Run(":8080")
}