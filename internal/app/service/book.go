package service

import (
	"book-store/internal/app/dto"
	"book-store/internal/app/repository"
)

type Book interface {
	Get(id int) (dto.Book, error)
}
type book struct {
	bookRepository repository.Book
}

func NewBook(bookRepository repository.Book) Book {
	return &book{
		bookRepository: bookRepository,
	}
}

func (book book) Get(id int) (dto.Book, error) {
	return book.bookRepository.FindById(id)
}
