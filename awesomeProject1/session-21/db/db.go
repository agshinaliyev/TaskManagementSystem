package db

import (
	"fmt"
	"github.com/agshinaliyev/ms-app/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

const (
	host     = "localhost"
	port     = "5432"
	username = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

var DB *gorm.DB

func Init() error {

	connStr := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s  sslmode=disable TimeZone=Asia/Baku",
		host, port, dbname, username, password)
	log.Println("Starting database connection: ", host, port)

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {

		log.Fatal("ActionLog.DbInit.error", err)
		return err
	}

	DB = db
	err = Migrate()
	if err != nil {
		return err
	}

	return nil
}

func GetDb() *gorm.DB {
	return DB
}

func Migrate() error {
	log.Println("Migration database started", host, port)
	err := GetDb().AutoMigrate(&model.Rate{})
	if err != nil {
		log.Fatal("ActionLog.Migrate.error", err)
		return err
	}
	return nil

}
