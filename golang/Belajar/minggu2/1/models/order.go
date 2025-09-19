package models

type Order struct {
	ID uint `json: "id" gorm: "primaryKey"`
	Item string `json: "item"`
	UserID uint `json: "user_id"`
}