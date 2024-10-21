package response

import "github.com/gin-gonic/gin"

// 响应体包装类
type ResultData struct {
	Code      int         `json:"code"`
	Data      interface{} `json:"data"`
	Msg       string      `json:"msg"`
	Total     int         `json:"total"`
	PageIndex int         `json:"pageIndex"`
	PageSize  int         `json:"pageSize"`
}

func Success(gin *gin.Context) {
	SuccessReturn(gin, nil)
}

func SuccessReturn(gin *gin.Context, data interface{}) {
	gin.JSON(200, ResultData{Code: 0, Data: data, Msg: "成功"})
}

func Failed(gin *gin.Context, error *ErrorCode) {
	if error == nil {
		error = &DEFAULT_ERROR
	}
	gin.JSON(500, ResultData{Code: 0, Msg: error.Msg})
}
