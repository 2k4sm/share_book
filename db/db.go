package db

import "time"

type Book struct {
	book_id  int `gorm:"primaryKey"`
	name     string
	author   string
	isbn     int `gorm:"unique"`
	yor      time.Time
	owner_id int
}

type Borrower struct {
	book_id           int `gorm:"foreignKey"`
	borrow_id         int `gorm:"primaryKey"`
	borrow_start_time time.Time
	borrow_end_time   time.Time
}
