package utils

import "github.com/gin-gonic/gin"

func TrueResponse(c *gin.Context, msg string, data any) {
	c.JSON(200, gin.H{
		"status": true,
		"msg":    msg,
		"data":   data,
	},
	)
}

func FalseResponse(c *gin.Context, msg string) {
	c.JSON(200, gin.H{
		"status": false,
		"msg":    msg,
		"data":   nil,
	},
	)
}
