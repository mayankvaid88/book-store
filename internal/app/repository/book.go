package repository

import (
	"book-store/internal/app/dto"
	"database/sql"
	"errors"
	"log"
)

const ()

type Book interface {
	Save(book dto.Book)
	FindById(id int) (dto.Book, error)
}

type book struct {
	*sql.DB
}

func NewBook(db *sql.DB) Book {
	return &book{
		db,
	}
}

func (book book) Save(bookDto dto.Book) {
	panic("yet to implement")
}

func (book book) FindById(bookId int) (dto.Book, error) {
	var bookObj dto.Book
	query := book.DB.QueryRow("select id, name, author, price from book b where b.id=$1", bookId)
	if err := query.Scan(&bookObj.Id, &bookObj.Name, &bookObj.Author, &bookObj.Price); err != nil {
		log.Print("error while scanning database ", err)
		return dto.Book{}, errors.New("unable to fetch data from database")
	}
	return bookObj, nil
}
