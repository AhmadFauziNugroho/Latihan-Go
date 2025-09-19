package main

import (
	"fmt"
	"1/database"
	"1/models"
)

func main() {
	database.Connect()
	db := database.DB

	db.AutoMigrate(&models.User{}, &models.Order{})

	user := models.User{Name: "Sien"}
	db.Create(&user)

	orders := []models.Order{
		{Item: "Laptop", UserID: user.ID},
		{Item: "Mouse", UserID: user.ID},
	}
	db.Create(&orders)


	var u models.User
	db.Preload("Orders").First(&u, user.ID)

	fmt.Println("User:", u.Name)
	for _, o := range u.Orders {
		fmt.Println("Order:", o.Item)
	}
}