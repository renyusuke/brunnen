package helper

import (
	"github.com/gin-gonic/gin"
)

func Res(c *gin.Context, httpCode int, data any, message any) {
	c.JSON(
		httpCode,
		gin.H{
			"message": message,
			"data":    data,
		},
	)
}
