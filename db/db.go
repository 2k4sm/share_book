package db

import "time"

// Book database Model.
type Book struct {
	Book_id int       `gorm:"primaryKey; not null" json:"book_id"`
	Name    string    `json:"name" gorm:"not null"`
	Author  string    `json:"author" gorm:"not null"`
	ISBN    int       `json:"isbn" gorm:"not null; unique"`
	AddedOn time.Time `json:"addedon"`
}

// Borrower database Model.
type Borrower struct {
	Book_id      uint      `gorm:"foreignKey; not null" json:"book_id"`
	Borrow_id    uint      `gorm:"primaryKey; not null" json:"borrow_id"`
	Borrow_start time.Time `json:"borrow_start" gorm:"not null"`
	Borrow_end   time.Time `json:"borrow_end" gorm:"not null"`
}
