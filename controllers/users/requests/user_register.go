package requests

import "final_project/business/users"

type UserRegister struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Phone_number int    `json:"phone_number"`
	Gender       string `json:"gender"`
	Dob          string `json:"dob"`
}

func (user *UserRegister) ToDomain() users.Domain {
	return users.Domain{
		Name:         user.Name,
		Email:        user.Email,
		Password:     user.Password,
		Phone_number: user.Phone_number,
		Gender:       user.Gender,
		Dob:          user.Dob,
	}
}
