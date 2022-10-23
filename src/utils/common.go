package utils

import "github.com/gin-gonic/gin"

type Response struct {
	Status bool        `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}

func TrueResponse(c *gin.Context, msg string, data any) {
	c.JSON(200, Response{
		Status: true,
		Msg:    msg,
		Data:   data,
	})
}

func FalseResponse(c *gin.Context, msg string) {
	c.JSON(200, Response{
		Status: false,
		Msg:    msg,
		Data:   nil,
	})
}
