package util

import "github.com/gin-gonic/gin"

type Gin struct {
	Ctx *gin.Context
}

type Response struct {
	Code    string      `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

func (g *Gin) Response(code string, msg string, data interface{}) {
	g.Ctx.JSON(200, Response{
		Code:    code,
		Message: msg,
		Data:    data,
	})
}
