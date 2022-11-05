package utils

import (
	"os"

	"github.com/gin-gonic/gin"
)

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

// WriteFile 写文件
func WriteFile(filename string, content []byte) (err error) {
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	if _, err = f.Write(content); err != nil {
		return err
	}
	return nil
}
