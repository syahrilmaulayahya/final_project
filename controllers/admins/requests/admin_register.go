package requests

import "final_project/business/admins"

type AdminRegister struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (admin *AdminRegister) ToDomain() admins.AdminDomain {
	return admins.AdminDomain{
		Name:     admin.Name,
		Email:    admin.Email,
		Password: admin.Password,
	}
}
