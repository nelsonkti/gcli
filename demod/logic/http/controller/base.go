package controller

import (
	"github.com/labstack/echo/v4"
)

//分页请求
type PageRequest struct {
	Page    int `form:"page" json:"page"`
	PerPage int `form:"per_page" json:"per_page"`
}

//分页返回
type PageResponse struct {
	PageRequest
	Total int         `json:"total"`
	Data  interface{} `json:"data"`
}

type qiniuRequest struct {
	Bucket string `form:"bucket" json:"bucket"`
	Key    string `form:"key" json:"key"`
}

func NewPageRequest(c echo.Context) PageRequest {
	var request PageRequest
	_ = c.Bind(&request)
	return DefaultPageRequest(request)
}

func DefaultPageRequest(request PageRequest) PageRequest {
	if request.Page < 1 {
		request.Page = 1
	}
	if request.PerPage < 1 {
		request.PerPage = 15
	}
	return request
}

func NewPageResponse(p PageRequest) PageResponse {
	return PageResponse{PageRequest: p}
}
