package utils

import (
	"log"

	"github.com/gin-gonic/gin"
)

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ShowResult(c *gin.Context, status int, data any, message string) {
	c.JSON(status, map[string]any{

		"message": message,
		"status":  status,
		"data":    data,
	})
}
