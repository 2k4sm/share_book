package handlers

import (
	"log"

	"github.com/2k4sm/share_book/db"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var book_db *gorm.DB
var borrower_db *gorm.DB

func init() {
	// Create the dbs and automigrate them.

	book_db, err := gorm.Open(sqlite.Open("books.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect book_db due to error:", err)
	}
	book_db.AutoMigrate(&db.Book{})

	borrower_db, err = gorm.Open(sqlite.Open("borrower.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to borrower_db due to error:", err)
	}
}
