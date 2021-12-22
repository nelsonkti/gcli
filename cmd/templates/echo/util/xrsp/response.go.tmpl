package xrsp

import "net/http"

const (
	CodeSuccess = 0 // 成功返回
	CodeFail    = 1
)

type PaginateData struct {
	Page    int         `json:"page"`
	PerPage int         `json:"per_page"`
	Total   int64       `json:"total"`
	Data    interface{} `json:"data"`
}

type RespData struct {
	Status  int         `json:"status"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Response *RespData

func Error(err error, code ...int) *RespData {
	codeInt := CodeFail

	status := http.StatusBadRequest
	if len(code) > 0 {
		codeInt = code[0]
		status = code[0]
	}

	return &RespData{
		Status:  status,
		Code:    codeInt,
		Message: err.Error(),
	}
}

func ErrorText(text string, code ...int) *RespData {
	codeInt := CodeFail

	status := http.StatusBadRequest
	if len(code) > 0 {
		codeInt = code[0]
		status = code[0]
	}

	return &RespData{
		Status:  status,
		Code:    codeInt,
		Message: text,
	}
}

func Paginate(total int64, page int, perPage int, data interface{}) *RespData {
	return &RespData{
		Code:    CodeSuccess,
		Status:  http.StatusOK,
		Message: "成功",
		Data: &PaginateData{
			PerPage: perPage,
			Page:    page,
			Total:   total,
			Data:    data,
		},
	}
}

func Data(data interface{}) *RespData {
	return &RespData{
		Code:    CodeSuccess,
		Status:  http.StatusOK,
		Data:    data,
		Message: "成功",
	}
}

func Nil() *RespData {
	return &RespData{
		Code:    CodeSuccess,
		Status:  http.StatusOK,
		Message: "成功",
	}
}
