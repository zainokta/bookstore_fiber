package controllers

import "github.com/zainokta/bookstore_fiber/models"

type BookHandler struct {
	bookRepo models.BookRepository
}

func NewBookHandler(bookRepo models.BookRepository) *BookHandler {
	return &BookHandler{
		bookRepo: bookRepo,
	}
}
