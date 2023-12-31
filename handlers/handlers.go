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

// View all the books that are currently in share.
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

// Borrow a book for a certain Period of time(for now it's a week).
func BorrowBook(ctx *fiber.Ctx) error {
	ctx.Response().Header.SetContentType("application/json")

	borrower_db, err = gorm.Open(sqlite.Open("borrower.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Error Occured while connecting to borrower.db:", err)
	}

	book_db, err = gorm.Open(sqlite.Open("books.db"), &gorm.Config{})

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
	isBookAvailable := false
	for _, book := range books {
		if book.Book_id == bookIdInt {

			isBookAvailable = true
			newBorrow.Book_id = uint(bookIdInt)
			newBorrow.Name = book.Name
			newBorrow.AddedOn = book.AddedOn
			newBorrow.Author = book.Author
			newBorrow.ISBN = book.ISBN

			newBorrow.Borrow_start = time.Now()
			newBorrow.Borrow_end = time.Date(time.Now().Year(), time.Now().Month(), (time.Now().Day() + 7), 0, 0, 0, 0, time.Now().Location())
		}
	}

	if !isBookAvailable {
		return ctx.JSON(fiber.Map{
			"available_books": books,
			"error":           "book not available for borrowing.",
		})
	}

	book_db.Delete(&db.Book{}, book_id)

	book_db.Order("name ASC").Find(&books)

	borrower_db.Create(newBorrow)
	borrowedBooks := []*db.Borrower{}

	borrower_db.Order("book_id ASC").Find(&borrowedBooks)

	return ctx.JSON(fiber.Map{
		"booksinshare":  books,
		"borrowedbooks": borrowedBooks,
	})

}

// View all the borrowed books.
func ViewBorrowedBooks(ctx *fiber.Ctx) error {
	ctx.Response().Header.SetContentType("application/json")

	borrower_db, err = gorm.Open(sqlite.Open("borrower.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Error Occured while connecting to borrower.db:", err)
	}

	borrowedBooks := []*db.Borrower{}

	borrower_db.Order("book_id ASC").Find(&borrowedBooks)

	return ctx.JSON(borrowedBooks)

}

// Returns a borrowed book.
func ReturnBorrowedBook(ctx *fiber.Ctx) error {
	ctx.Response().Header.SetContentType("application/json")

	book_id := ctx.Params("bookid")
	bookIdInt, err := strconv.Atoi(book_id)
	if err != nil {
		log.Fatal("error converting book_id to int: ", err)
	}

	borrow_id := ctx.Params("borrowid")
	borrowIdInt, err := strconv.Atoi(borrow_id)
	if err != nil {
		log.Fatal("error converting borrow_id to int: ", err)
	}

	borrower_db, err = gorm.Open(sqlite.Open("borrower.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Error Occured while connecting to borrower.db:", err)
	}

	book_db, err = gorm.Open(sqlite.Open("books.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Error Occured while connecting to books.db:", err)
	}

	allBorrowedBooks := []*db.Borrower{}
	borrower_db.Order("borrow_id ASC").Find(&allBorrowedBooks)

	bookToaddBack := new(db.Book)

	isAvailable := false
	for _, borrowedBook := range allBorrowedBooks {
		if borrowedBook.Book_id == uint(bookIdInt) && borrowedBook.Borrow_id == uint(borrowIdInt) {
			isAvailable = true
			bookToaddBack.Name = borrowedBook.Name
			bookToaddBack.Author = borrowedBook.Author
			bookToaddBack.ISBN = borrowedBook.ISBN
			bookToaddBack.AddedOn = time.Now()

		}
	}

	if !isAvailable {
		return ctx.JSON(fiber.Map{
			"error":         "book not borrowed.",
			"borrowedBooks": allBorrowedBooks,
		})
	}

	borrower_db.Delete(&db.Borrower{}, borrow_id)

	borrower_db.Order("name ASC").Find(&allBorrowedBooks)

	book_db.Create(bookToaddBack)
	booksAvailable := []*db.Book{}

	book_db.Order("name ASC").Find(&booksAvailable)

	return ctx.JSON(fiber.Map{
		"booksinshare":  booksAvailable,
		"borrowedbooks": allBorrowedBooks,
	})

}
