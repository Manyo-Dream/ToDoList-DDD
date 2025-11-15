package types

import "github.com/manyodream/todolist-ddd/consts"

type Response struct {
	Status  int    `json:"status"`
	Data    any    `json:"data"`
	Msg     string `json:"msg"`
	Error   string `json:"error"`
	TraceID string `json:"trace_id"`
}

func RespSuccess(code ...int) *Response {
	status := consts.Success
	if len(code) > 0 {
		status = code[0]
	}
	return &Response{Status: status,
		Data: consts.GetMsg(status),
		Msg:  consts.GetMsg(status)}
}

func RespSuccessWithData(data any, code ...int) *Response {
	status := consts.Success
	if len(code) > 0 {
		status = code[0]
	}
	return &Response{Status: status,
		Data: data,
		Msg:  consts.GetMsg(status)}
}

func RespError(err error, data string, code ...int) *Response {
	status := consts.Error
	if len(code) > 0 {
		status = code[0]
	}
	return &Response{Status: status,
		Data:  data,
		Msg:   consts.GetMsg(status),
		Error: err.Error()}
}
