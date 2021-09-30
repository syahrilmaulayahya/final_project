package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type BaseRespons struct {
	Meta struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	} `json:"meta"`
	Data interface{} `json:"data"`
}

func NewSuccessResponse(c echo.Context, data interface{}) error {
	response := BaseRespons{}
	response.Meta.Status = http.StatusOK
	response.Meta.Message = "success"
	response.Data = data
	return c.JSON(http.StatusOK, response)
}

func NewErrorResponse(c echo.Context, status int, err error) error {
	response := BaseRespons{}
	response.Meta.Status = status
	response.Meta.Message = err.Error()
	response.Data = nil
	return c.JSON(status, response)
}

type BaseResponsDetails struct {
	Meta struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	} `json:"meta"`
	Data1 interface{} `json:"data1"`
	Data2 interface{} `json:"data2"`
	Data3 interface{} `json:"data3"`
}

func NewSuccessResponseDetails(c echo.Context, data1, data2, data3 interface{}) error {
	response := BaseResponsDetails{}
	response.Meta.Status = http.StatusOK
	response.Meta.Message = "success"
	response.Data1 = data1
	response.Data2 = data2
	response.Data3 = data3
	return c.JSON(http.StatusOK, response)
}
