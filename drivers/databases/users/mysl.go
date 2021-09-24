package users

import (
	"context"
	"final_project/business/users"

	"gorm.io/gorm"
)

type MysqlUserRepository struct {
	Conn *gorm.DB
}

func NewMysqlRepository(conn *gorm.DB) users.Repository {
	return &MysqlUserRepository{
		Conn: conn,
	}
}

func (rep MysqlUserRepository) Register(ctx context.Context, domain users.Domain) (users.Domain, error) {
	var user User
	user.Name = domain.Name
	user.Email = domain.Email
	user.Password = domain.Password
	user.Phone_number = domain.Phone_number
	user.Gender = domain.Gender
	user.Dob = domain.Dob
	result := rep.Conn.Create(&user)
	if result.Error != nil {
		return users.Domain{}, result.Error
	}
	return user.ToDomain(), nil

}
func (rep MysqlUserRepository) Login(ctx context.Context, email, password string) (users.Domain, error) {
	var user User
	result := rep.Conn.First(&user, "email = ? AND password = ?",
		email, password)

	if result.Error != nil {
		return users.Domain{}, result.Error
	}
	return user.ToDomain(), nil
}

func (rep MysqlUserRepository) Get(ctx context.Context) (users.Domain, error) {
	var user User
	result := rep.Conn.Find(&user)
	if result.Error != nil {
		return users.Domain{}, result.Error
	}
	return user.ToDomain(), nil
}
