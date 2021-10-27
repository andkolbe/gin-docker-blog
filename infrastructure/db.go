package infrastructure

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// db struct
type Database struct {
	DB *gorm.DB
}

// New Database: initializes and returns mysql db
func NewDatabase() Database {
	USER := os.Getenv("DB_USER")
    PASS := os.Getenv("DB_PASSWORD")
    HOST := os.Getenv("DB_HOST")
    DBNAME := os.Getenv("DB_NAME")

    URL := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS, HOST, DBNAME)
    fmt.Println(URL)
	// a new mysql connection is opened with the gorm.Open method from URL
	db, err := gorm.Open(mysql.Open(URL))

	if err != nil {
		panic("Failed to connect to database!")
	}

	fmt.Println("Database connect established")
	// Database struct is returned with gorm database instance as parameter, which the application uses to access the db
	return Database{
		DB: db,
	}
}