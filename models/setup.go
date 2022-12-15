package models

import (
	"gorm.io/driver/postgres"
  "gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(postgres.New(postgres.Config{
		DSN: "host=localhost user=dangutt dbname=asteroids_dev port=5432 sslmode=disable", // data source name, refer https://github.com/jackc/pgx
		PreferSimpleProtocol: true, // disables implicit prepared statement usage. By default pgx automatically uses the extended protocol
	}), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	err = database.AutoMigrate(&Game{},&Player{})
	if err != nil {
		return
	}

	DB = database
}