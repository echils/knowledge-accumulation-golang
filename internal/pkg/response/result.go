package response

// 响应体包装类
type ResultData struct {
	Code      int         `json:"code"`
	Data      interface{} `json:"data"`
	Msg       string      `json:"msg"`
	Total     int         `json:"total"`
	PageIndex int         `json:"pageIndex"`
	PageSize  int         `json:"pageSize"`
}

func Success() (result ResultData) {
	return ResultData{Code: 0, Data: nil, Msg: "成功"}
}
func SuccessReturn(data interface{}) (result ResultData) {
	return ResultData{Code: 0, Data: data, Msg: "成功"}
}
func Failed(error *ErrorCode) (result ResultData) {
	if error == nil {
		error = &DEFAULT_ERROR
	}
	return ResultData{Code: error.Code, Msg: error.Msg}
}
