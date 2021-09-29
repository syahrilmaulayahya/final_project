package helpers

import (
	"net/mail"
	"unicode"
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
