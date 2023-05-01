package helper

import (
	"github.com/gin-gonic/gin"
)

func Res(c *gin.Context, httpCode int, data any, message string) {
	c.JSON(
		httpCode,
		gin.H{
			"message": message,
			"data":    data,
		},
	)
}
