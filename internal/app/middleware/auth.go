package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		log.Print("inside authentication middleware")
		authToken := context.Request.Header.Get("Authentication")
		log.Print("auth token lenth", len(authToken))
		if len(authToken) == 0 {
			log.Print("inside if block")
			context.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		context.Next()
		log.Print("After context next in Auth")
	}
}

func Logging() gin.HandlerFunc {
	return func(context *gin.Context) {
		log.Print("logging middleware kicked in ")
		context.Next()
		log.Print("After context next in Logging")
	}
}
