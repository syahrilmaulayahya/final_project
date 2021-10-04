package admins

import (
	"context"
	"final_project/business/admins"

	"gorm.io/gorm"
)

type MysqlAdminRepository struct {
	Conn *gorm.DB
}

func NewMysqlRepository(conn *gorm.DB) admins.Repository {
	return &MysqlAdminRepository{
		Conn: conn,
	}
}

func (rep MysqlAdminRepository) Register(ctx context.Context, domain admins.AdminDomain) (admins.AdminDomain, error) {
	var admin Admin
	admin.Name = domain.Name
	admin.Email = domain.Email
	admin.Password = domain.Password
	createAdmin := rep.Conn.Create(&admin)
	if createAdmin.Error != nil {
		return admins.AdminDomain{}, createAdmin.Error
	}
	return admin.ToDomain(), nil
}

func (rep MysqlAdminRepository) Login(ctx context.Context, email, password string) (admins.AdminDomain, error) {
	var admin Admin

	result := rep.Conn.First(&admin, "email = ? AND password = ?", email, password)
	if result.Error != nil {
		return admins.AdminDomain{}, result.Error
	}
	return admin.ToDomain(), nil
}
