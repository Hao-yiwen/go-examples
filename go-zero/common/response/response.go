package response

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-user/common/errorx"
)

type Body struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func Response(w http.ResponseWriter, resp interface{}, err error) {
	var body Body

	if err != nil {
		switch e := err.(type) {
		case *errorx.CodeError:
			body.Code = e.GetCode()
			body.Msg = e.GetMsg()
		default:
			body.Code = errorx.CodeInternal
			body.Msg = err.Error()
		}
	} else {
		body.Code = errorx.CodeSuccess
		body.Msg = "success"
		body.Data = resp
	}

	httpx.OkJson(w, body)
}

func Success(w http.ResponseWriter, data interface{}) {
	Response(w, data, nil)
}

func Error(w http.ResponseWriter, err error) {
	Response(w, nil, err)
}

func ErrorCode(w http.ResponseWriter, code int) {
	Response(w, nil, errorx.NewCodeError(code))
}

func ErrorMsg(w http.ResponseWriter, code int, msg string) {
	Response(w, nil, errorx.NewCodeErrorMsg(code, msg))
}
