package middlewares

import (
	"final_project/constans"
	"time"

	"github.com/golang-jwt/jwt"
)

type JwtCustomClaims struct {
	ID int `json:"id"`
	jwt.StandardClaims
}

func GenerateTokenJWT(id int) (string, error) {
	claims := JwtCustomClaims{
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * 1).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwtToken, err := token.SignedString([]byte(constans.SECRET_JWT))
	if err != nil {
		return "", nil
	}

	return jwtToken, nil
}
