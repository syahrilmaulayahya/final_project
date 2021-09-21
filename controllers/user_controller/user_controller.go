package user_controller

import (
	"errors"
	"final_project/configs"
	"final_project/models/responses"
	"final_project/models/users"
	"net/http"
	"net/mail"

	"github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// source https://stackoverflow.com/questions/66624011/how-to-validate-an-email-address-in-go
func CheckEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
func UserRegisterController(c echo.Context) error {
	var userRegister users.UserRegister
	var userDB users.User

	c.Bind(&userRegister)

	switch {
	case userRegister.Username == "":
		return c.JSON(http.StatusBadRequest, responses.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "username masih kosong",
			Data:    nil,
		})
	case userRegister.Name == "":
		return c.JSON(http.StatusBadRequest, responses.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Nama masih kosong",
			Data:    nil,
		})
	case userRegister.Email == "":
		return c.JSON(http.StatusBadRequest, responses.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Email masih kosong",
			Data:    nil,
		})
	case userRegister.Password == "":
		return c.JSON(http.StatusBadRequest, responses.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Password masih kosong",
			Data:    nil,
		})
	case CheckEmail(userRegister.Email) == false:
		return c.JSON(http.StatusBadRequest, responses.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Format Email salah",
			Data:    nil,
		})

	}
	userDB.Username = userRegister.Username
	userDB.Name = userRegister.Name
	userDB.Email = userRegister.Email
	userDB.Password = userRegister.Password

	result := configs.DB.Create(&userDB)

	//source https://github.com/go-gorm/gorm/issues/4037
	var mysqlErr *mysql.MySQLError
	if result.Error != nil {
		if errors.As(result.Error, &mysqlErr) && mysqlErr.Number == 1062 {
			return c.JSON(http.StatusBadRequest, responses.BaseResponse{
				Code:    http.StatusBadRequest,
				Message: "Email sudah digunakan",
				Data:    nil,
			})
		} else {
			return c.JSON(http.StatusInternalServerError, responses.BaseResponse{
				Code:    http.StatusInternalServerError,
				Message: "Error ketika input data user ke DB",
				Data:    nil,
			})
		}

	}

	return c.JSON(http.StatusOK, responses.BaseResponse{
		Code:    http.StatusOK,
		Message: "registrasi berhasil",
		Data:    userDB,
	})
}

func GetUserController(c echo.Context) error {
	users := []users.User{}
	result := configs.DB.Find(&users)

	if result.Error != nil {
		if result.Error != gorm.ErrRecordNotFound {
			return c.JSON(http.StatusInternalServerError, responses.BaseResponse{
				Code:    http.StatusInternalServerError,
				Message: "Error ketika mendapatkan data user dari DB",
				Data:    nil,
			})

		}
	}
	return c.JSON(http.StatusOK, responses.BaseResponse{
		Code:    http.StatusOK,
		Message: "Berhasil mendapatkan data user dari DB",
		Data:    users,
	})

}
