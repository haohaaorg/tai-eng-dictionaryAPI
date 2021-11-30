package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open("sqlite3", "./tai-eng.db")
	if err != nil {
		panic("failed to connect database")
	}
	DB = db
}

func CloseDatabase() {
	DB.Close()
}
