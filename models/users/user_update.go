package users

type NewUserData struct {
	Username     string `json:"username"`
	Picture_url  string `json:"picture_url"`
	Phone_number int    `json:"phone_number"`
	Email        string `gorm:"unique" json:"email"`
	Password     string `json:"-"`
	Name         string `json:"name"`
	Dob          string `json:"dob"`
}
