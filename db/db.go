package db

import "time"

// Book database Model.
type Book struct {
	Book_id  int `gorm:"primaryKey"`
	Name     string
	Author   string
	ISBN     int `gorm:"unique"`
	YOR      time.Time
	Owner_id int
}

// Borrower database Model.
type Borrower struct {
	Book_id           int `gorm:"foreignKey"`
	Borrow_id         int `gorm:"primaryKey"`
	Borrow_start_time time.Time
	Borrow_end_time   time.Time
}
