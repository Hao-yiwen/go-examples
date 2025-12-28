package errorx

// 错误码定义
const (
	// 通用错误码 1000-1999
	CodeSuccess      = 0
	CodeUnknown      = 1000
	CodeInvalidParam = 1001
	CodeUnauthorized = 1002
	CodeForbidden    = 1003
	CodeNotFound     = 1004
	CodeInternal     = 1005

	// 用户相关错误码 2000-2999
	CodeUserNotFound      = 2001
	CodeUserAlreadyExists = 2002
	CodePasswordError     = 2003
	CodeUserDisabled      = 2004

	// 认证相关错误码 3000-3999
	CodeTokenInvalid  = 3001
	CodeTokenExpired  = 3002
	CodeTokenGenerate = 3003

	// 角色相关错误码 4000-4999
	CodeRoleNotFound      = 4001
	CodeRoleAlreadyExists = 4002
)

var codeMsg = map[int]string{
	CodeSuccess:           "success",
	CodeUnknown:           "未知错误",
	CodeInvalidParam:      "参数错误",
	CodeUnauthorized:      "未授权",
	CodeForbidden:         "禁止访问",
	CodeNotFound:          "资源不存在",
	CodeInternal:          "内部错误",
	CodeUserNotFound:      "用户不存在",
	CodeUserAlreadyExists: "用户已存在",
	CodePasswordError:     "密码错误",
	CodeUserDisabled:      "用户已禁用",
	CodeTokenInvalid:      "Token无效",
	CodeTokenExpired:      "Token已过期",
	CodeTokenGenerate:     "Token生成失败",
	CodeRoleNotFound:      "角色不存在",
	CodeRoleAlreadyExists: "角色已存在",
}

type CodeError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func NewCodeError(code int) *CodeError {
	return &CodeError{
		Code: code,
		Msg:  codeMsg[code],
	}
}

func NewCodeErrorMsg(code int, msg string) *CodeError {
	return &CodeError{
		Code: code,
		Msg:  msg,
	}
}

func (e *CodeError) Error() string {
	return e.Msg
}

func (e *CodeError) GetCode() int {
	return e.Code
}

func (e *CodeError) GetMsg() string {
	return e.Msg
}
