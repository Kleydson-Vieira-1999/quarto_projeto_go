package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
	err error
)

func ConnectDB() {
	dsn := "host=localhost port=5432 user=doctor dbname=dev_db sslmode=disable"
	
	DB, err = gorm.Open(postgres.Open(dsn))

	if err != nil {
		log.Panic("erro ao conectar ao banco de dados", err.Error())
	}
}
