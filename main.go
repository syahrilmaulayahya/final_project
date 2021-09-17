package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Users struct {
	Id           int    `json:"id"`
	Username     string `json:"username"`
	Picture_url  string `json:"picture_url"`
	Phone_number int    `json:"phone_number"`
	Email        string `json:"email"`
	// Password     string `json:"password"`
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Dob    string `json:"dob"`
}

type BaseResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func main() {
	e := echo.New()
	e.GET("/v1/users", GetUserController)
	e.Start(":8000")
}

func GetUserController(c echo.Context) error {
	user := Users{1, "puppy", "https://id.depositphotos.com/stock-photos/puppy.html", +62856414411299, "puppy@gmail.com", "pupy", "male", "2000-3-1"}
	return c.JSON(http.StatusOK, BaseResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    user,
	})
}
