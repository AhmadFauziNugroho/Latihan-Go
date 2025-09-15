package database

import (
	"log"

	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)

var DB *gorm.DB

func Connect() {
	var err error
	DB, err = gorm.Open(sqlite.Open("app.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("gagal konek database:	 %v", err)
	}
}