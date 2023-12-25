package handlers

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/2k4sm/share_book/db"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Initiating the databases...

// book database.
var book_db *gorm.DB

// borrower database.
var borrower_db *gorm.DB
var err error

func init() {
	// Initiate the dbs and automigrate them.

	book_db, err = gorm.Open(sqlite.Open("books.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Error Occured while connecting to books.db:", err)
	}

	err = book_db.AutoMigrate(&db.Book{})

	if err != nil {
		log.Default().Println("Error Occured while automigrating database:,", err)
	}

	borrower_db, err = gorm.Open(sqlite.Open("borrower.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Error Occured while connecting to borrower.db:", err)
	}

	err = borrower_db.AutoMigrate(&db.Borrower{})

	if err != nil {
		log.Default().Println("Error Occured while automigrating database:,", err)
	}
}

// Handlers...

// Share a book for others to view and borrow.
func ShareBook(ctx *fiber.Ctx) error {
	ctx.Response().Header.SetContentType("application/json")
	book_db, err := gorm.Open(sqlite.Open("books.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Error Occured while connecting to books.db:", err)
	}

	// This creates a new book struct object to parse the request body.
	newbook := new(db.Book)

	if err := ctx.BodyParser(newbook); err != nil {
		return fmt.Errorf("error occured while sharing book: %v", err)
	}

	if newbook.Name == "" || newbook.Author == "" || newbook.ISBN == 0 {
		return fmt.Errorf("error while creating dbcalls:one of the required Params is empty")
	}
	newbook.AddedOn = time.Now()

	book_db.Create(newbook)

	books := []*db.Book{}

	book_db.Order("Name ASC").Find(&books)

	return ctx.JSON(books)
}

// View shared books.
func ViewSharedBooks(ctx *fiber.Ctx) error {
	ctx.Response().Header.SetContentType("application/json")
	book_db, err := gorm.Open(sqlite.Open("books.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Error Occured while connecting to books.db:", err)
	}

	books := []*db.Book{}

	book_db.Order("Name ASC").Find(&books)

	return ctx.JSON(books)
}

// Borrow a book for a certain duration of time.
func BorrowBook(ctx *fiber.Ctx) error {
	borrower_db, err = gorm.Open(sqlite.Open("borrower.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Error Occured while connecting to borrower.db:", err)
	}

	book_db, err := gorm.Open(sqlite.Open("books.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Error Occured while connecting to books.db:", err)
	}

	book_id := ctx.Params("bookid")
	bookIdInt, err := strconv.Atoi(book_id)

	if err != nil {
		log.Fatal("error converting book_id to int: ", err)
	}

	books := []*db.Book{}

	book_db.Order("name ASC").Find(&books)

	//Creates a new Borrow if the book with bookId is present.
	newBorrow := new(db.Borrower)
	for _, book := range books {
		if book.Book_id == bookIdInt {
			newBorrow.Book_id = uint(bookIdInt)
			newBorrow.Borrow_start = time.Now()
			newBorrow.Borrow_end = time.Date(time.Now().Year(), time.Now().Month(), (time.Now().Day() + 7), 0, 0, 0, 0, nil)
		}
	}
	book_db.Delete(bookIdInt)

	book_db.Order("name ASC").Find(&books)

}
