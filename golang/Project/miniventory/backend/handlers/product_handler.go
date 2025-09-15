package handlers

import (
	"miniventory/backend/database"
	"miniventory/backend/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllProducts(c *gin.Context) {
	var items []models.Product
	database.DB.Find(&items)
	c.JSON(http.StatusOK, items)
}

func CreateProduct(c *gin.Context) {
	var body models.Product
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}
	database.DB.Create(&body)
	c.JSON(http.StatusCreated,body)
}

func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	pid, _ := strconv.Atoi(id)

	var existing models.Product
	if err := database.DB.First(&existing, pid).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product tidak ada"})
		return
	}

	var body models.Product
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	existing.Name = body.Name
	existing.Price = body.Price
	existing.Stock = body.Stock
	database.DB.Save(&existing)

	c.JSON(http.StatusOK, existing)
}

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	pid, _ := strconv.Atoi(id)
	database.DB.Delete(&models.Product{}, pid)
	c.Status(http.StatusNoContent)
}