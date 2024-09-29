package response

// 错误码
type ErrorCode struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

var (
	//默认全局异常
	DEFAULT_ERROR = ErrorCode{Code: -1, Msg: "响应异常"}
)
