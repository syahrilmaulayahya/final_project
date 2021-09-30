package helpers

import (
	"net/mail"
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

func CheckEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
func CheckPassword(s string) bool {
	var (
		hasMinLen = false
		hasUpper  = false
		hasLower  = false
		hasNumber = false
	)
	if len(s) >= 7 {
		hasMinLen = true
	}
	for _, char := range s {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true

		}
	}
	return hasMinLen && hasUpper && hasLower && hasNumber
}

func Hash(password string) (string, error) {
	result, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(result), nil
}
