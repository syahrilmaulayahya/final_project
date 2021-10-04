package admins

import (
	"final_project/business/admins"
	"time"
)

type Admin struct {
	ID        int `gorm:"primaryKey"`
	Name      string
	Email     string `gorm:"unique"`
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (admin *Admin) ToDomain() admins.AdminDomain {
	return admins.AdminDomain{
		ID:        admin.ID,
		Name:      admin.Name,
		Email:     admin.Email,
		Password:  admin.Password,
		CreatedAt: admin.CreatedAt,
		UpdatedAt: admin.UpdatedAt,
	}
}

func FromDomain(domain admins.AdminDomain) Admin {
	return Admin{
		ID:        domain.ID,
		Name:      domain.Name,
		Email:     domain.Email,
		Password:  domain.Password,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
