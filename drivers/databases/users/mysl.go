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
	user.Address = domain.Address
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

func (rep MysqlUserRepository) Details(ctx context.Context, id int) (users.Domain, error) {
	var user User
	result := rep.Conn.Find(&user, "id = ?", id)
	if result.Error != nil {
		return users.Domain{}, result.Error
	}
	return user.ToDomain(), nil
}

func (rep MysqlUserRepository) UploadReview(ctx context.Context, domain users.Review_RatingDomain, id int) (users.Review_RatingDomain, error) {
	var review Review_Rating
	review.Review = domain.Review
	review.Rating = domain.Rating
	review.UserID = id
	review.ProductID = domain.ProductID
	result := rep.Conn.Create(&review)
	if result.Error != nil {
		return users.Review_RatingDomain{}, result.Error
	}
	return review.ToDomain(), nil

}
