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

func Hash(secret string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(secret), bcrypt.MinCost)
	return string(bytes), err
}
