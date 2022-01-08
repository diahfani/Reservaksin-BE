package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type BaseResponse struct {
	Meta struct {
		Status   int      `json:"status"`
		Message  string   `json:"Message"`
		Messages []string `json:"messages,omitempty"`
	} `json:"meta"`
	Data interface{} `json:"data"`
}

func NewSuccessResponse(c echo.Context, data interface{}) error {
	res := BaseResponse{}
	res.Meta.Status = http.StatusOK
	res.Meta.Message = "success"
	res.Data = data

	return c.JSON(http.StatusOK, res)
}

func NewErrorResponse(c echo.Context, status int, err error) error {
	res := BaseResponse{}
	res.Meta.Status = status
	res.Meta.Message = "error"
	res.Meta.Messages = []string{err.Error()}

	return c.JSON(status, res)
}
