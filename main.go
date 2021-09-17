package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Users struct {
	UserId       int    `json:"userId"`
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
	ev1users := e.Group("v1/users")
	ev1users.GET("", GetUserController)
	ev1users.GET("/:userId", DetailUserController)
	e.Start(":8000")
}

func GetUserController(c echo.Context) error {
	user := Users{1, "puppy", "https://id.depositphotos.com/stock-photos/puppy.html", +62856414411299, "puppy@gmail.com", "pupy", "male", "2000-3-1"}
	response := BaseResponse{http.StatusOK, "Success", user}
	return c.JSON(http.StatusOK, response)
}

func DetailUserController(c echo.Context) error {
	userId, err := strconv.Atoi(c.Param("userId"))
	response := BaseResponse{http.StatusOK, "Success", Users{UserId: userId}}
	failresponse := BaseResponse{http.StatusInternalServerError, "Gagal konversi userId", nil}
	if err != nil {
		return c.JSON(http.StatusOK, failresponse)
	}
	return c.JSON(http.StatusOK, response)
}
