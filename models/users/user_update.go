package users

type UserUpdate struct {
	Username     string `json:"username"`
	Picture_url  string `json:"picture_url"`
	Phone_number int    `json:"phone_number"`
	Password     string `json:"-"`
	Dob          string `json:"dob"`
}
