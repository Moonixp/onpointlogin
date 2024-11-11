package error_handler

import (
	"log"

	"github.com/gin-gonic/gin"
)

var logger = log.Default()

func ErrorHandler(err error) {
	if err != nil {
		logger.Fatal(err.Error())
	}
}

func PanicHandler(err error) {
	if err != nil {
		panic(err.Error())
	}
}

// can't use this because main function must return
func HttpErrorHandler(err error, code int, ctx *gin.Context) {
	if err != nil {
		ctx.JSON(code, gin.H{"error": err.Error()})
		return
	}
}
