package models

import "time"

type Book struct {
	ID        uint64
	Title     string
	Author    string
	Price     string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

type BookRepository interface {
	All() ([]Book, error)
	Find(id int64) (Book, error)
	Create(book *Book) error
	Update(id int64, book *Book) error
	Delete(id int64) error
}
