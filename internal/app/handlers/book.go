package handlers

import (
	"book-store/internal/app/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type Book interface {
	Get(ctx *gin.Context)
}

type book struct {
	bookService service.Book
}

func NewBook(bookSevice service.Book) Book {
	return &book{bookService: bookSevice}
}

func (book book) Get(ctx *gin.Context) {
	log.Print("inside GET handler")
	param := ctx.Param("bookId")
	id, err := strconv.Atoi(param)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}
	bookByGivenId, err := book.bookService.Get(id)
	if err != nil {
		log.Print("error while invoking api ", err)
		ctx.Status(http.StatusInternalServerError)
		return
	}
	ctx.AbortWithStatusJSON(http.StatusOK, bookByGivenId)
}
