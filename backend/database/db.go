package database

import (
	"todo/database/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Setup(dsn string) {
  db, err := gorm.Open(postgres.Open(dsn)) 

  if err != nil {
    panic("Couldn't connect to database")
  }

  db.AutoMigrate(&models.Todo{})

  DB = db
}
