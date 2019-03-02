package defs

type Err struct {
	Error     string `json:"error"`
	ErrorCode string `json:"error_code"` //状态码
}

type ErrorResponse struct {
	HttpSc int
	Error  Err //前者
}

//声明并初始化
var (
	ErrorRequestBodyParseFailed = ErrorResponse{HttpSc: 400, Error: Err{Error: "Request body is not correct", ErrorCode: "001"}} //请求错误
	ErrorNotAuthUser            = ErrorResponse{HttpSc: 401, Error: Err{Error: "user authentication failed", ErrorCode: "002"}}  //用户验证不通过
	ErrorDBError                = ErrorResponse{HttpSc: 500, Error: Err{Error: "DB ops failed", ErrorCode: "003"}}
	ErrorInternalFaults         = ErrorResponse{HttpSc: 500, Error: Err{Error: "Internal service error", ErrorCode: "004"}}
)
