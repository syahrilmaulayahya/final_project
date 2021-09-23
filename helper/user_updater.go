package helpers

import (
	"final_project/configs"
	"final_project/models/users"
)

type UserModel struct {
}

func (userModel UserModel) Update(user *users.User) error {
	result := configs.DB.Save(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (UserModel UserModel) FindById(id int) (users.User, error) {
	var user users.User
	result := configs.DB.Where("id = ?", id).First(&user)
	if result.Error != nil {
		return users.User{}, result.Error
	}
	return user, nil
}
