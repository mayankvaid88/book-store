package routes

import (
	"book-store/internal/app/handlers"
	"book-store/internal/app/middleware"
	"book-store/internal/app/repository"
	"book-store/internal/app/service"
	"book-store/internal/db"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(engine *gin.Engine) {
	dbConn := db.CreateConnection()
	bookRepository := repository.NewBook(dbConn)
	bookService := service.NewBook(bookRepository)
	bookHandler := handlers.NewBook(bookService)
	engine.Use(middleware.Auth())
	engine.Use(middleware.Logging())
	group := engine.Group("/api/bs/v1")
	{
		group.GET("/books/:bookId", bookHandler.Get)
	}
}
