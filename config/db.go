package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectionDb() *gorm.DB {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed Load .env Variable")
	}

	dbhost := os.Getenv("DB_HOST")
	dbuser := os.Getenv("DB_USER")
	dbpass := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbuser, dbpass, dbhost, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed Conecction Database")
	}

	return db

	// db.AutoMigrate(&models.User{}, models.Book{})
}

func CloseConnectionDb(db *gorm.DB) {
	Closedb, err := db.DB()
	if err != nil {
		panic("Failed Close Connectin")
	}

	Closedb.Close()
}
