package user_controller

import (
	"errors"
	"final_project/configs"
	helpers "final_project/helper"
	"final_project/middlewares"

	"final_project/models/responses"
	"final_project/models/users"
	"net/http"

	"github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

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
	case !helpers.CheckEmail(userRegister.Email):
		return c.JSON(http.StatusBadRequest, responses.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Format Email salah",
			Data:    nil,
		})
	case !helpers.CheckPassword(userRegister.Password):
		return c.JSON(http.StatusBadRequest, responses.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Password harus lebih dari 7 karakter, mengandung angka dan huruf kapital",
			Data:    nil,
		})
	}
	userDB.Username = userRegister.Username
	userDB.Name = userRegister.Name
	userDB.Email = userRegister.Email
	userDB.Password = userRegister.Password

	result := configs.DB.Create(&userDB)

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

func LoginController(c echo.Context) error {

	userLogin := users.UserLogin{}
	c.Bind(&userLogin)

	user := users.User{}

	result := configs.DB.First(&user, "email = ? AND password = ?",
		userLogin.Email, userLogin.Password)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusForbidden, responses.BaseResponse{
				Code:    http.StatusForbidden,
				Message: "email atau password salah",
				Data:    nil,
			})
		} else {
			return c.JSON(http.StatusInternalServerError, responses.BaseResponse{
				Code:    http.StatusInternalServerError,
				Message: "Ada keselahan di server",
				Data:    nil,
			})
		}

	}

	token, err := middlewares.GenerateTokenJWT(int(user.ID))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Ada keselahan di server",
			Data:    nil,
		})
	}

	userResponse := users.UserResponse{
		ID:           user.ID,
		Username:     user.Username,
		Picture_url:  user.Picture_url,
		Phone_number: user.Phone_number,
		Email:        user.Email,
		Token:        token,
		Name:         user.Name,
		Gender:       user.Gender,
		Dob:          user.Dob,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
		DeletedAt:    user.DeletedAt,
	}
	return c.JSON(http.StatusOK, responses.BaseResponse{
		Code:    http.StatusOK,
		Message: "Berhasil Login",
		Data:    userResponse,
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
