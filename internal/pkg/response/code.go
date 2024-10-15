package response

// 错误码
type ErrorCode struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

var (
	//默认全局异常
	DEFAULT_ERROR = ErrorCode{Code: -1, Msg: "系统响应异常"}
	//非法参数
	INVALID_PARAM = ErrorCode{Code: -2, Msg: "参数不合理"}
)
