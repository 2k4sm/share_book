package main

import (
	"log"

	"github.com/2k4sm/share_book/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "share_book v0.0.1",
	})

	// Adds a new book for share.
	app.Put("/api/v1/booky", handlers.ShareBook)

	// Retrieves all the books in share.
	app.Get("/api/v1/booky", handlers.ViewSharedBooks)

	// // Borrows a book from the shared books for a specified time.
	// app.Put()

	// // Returns the shared book.
	// app.Post()

	log.Fatal(app.Listen(":8000"))
}
