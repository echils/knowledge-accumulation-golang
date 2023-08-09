package model

import "github.com/gin-gonic/gin"

type ResultData struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(200, ResultData{
		Code: 200,
		Data: data,
		Msg:  "成功",
	})
}

func InsertSuccess(c *gin.Context, data interface{}) {
	c.JSON(200, ResultData{
		Code: 201,
		Data: data,
		Msg:  "成功",
	})
}

func UpdateSuccess(c *gin.Context, data interface{}) {
	c.JSON(200, ResultData{
		Code: 201,
		Data: data,
		Msg:  "成功",
	})
}

func DeleteSuccess(c *gin.Context, data interface{}) {
	c.JSON(200, ResultData{
		Code: 204,
		Data: data,
		Msg:  "成功",
	})
}

func Failed(c *gin.Context, msg string) {
	c.JSON(200, ResultData{
		Code: 502,
		Data: nil,
		Msg:  msg,
	})
}
