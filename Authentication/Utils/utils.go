package utils

import "github.com/gin-gonic/gin"

func CheckError(err error) error {
	if err != nil {
		panic(err.Error())
		return err
	} else {
		return nil
	}

}

func Showresult(c *gin.Context, status int, data any, message string) {
	c.JSON(status, map[string]any{

		"message": message,
		"status":  status,
		"data":    data,
	})
}
