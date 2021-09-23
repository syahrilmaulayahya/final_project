package middlewares

import (
	"final_project/constans"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type JwtCustomClaims struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	jwt.StandardClaims
}

func GenerateTokenJWT(id int, name string) (string, error) {
	claims := JwtCustomClaims{
		id,
		name,
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

func GetClaimsUserId(c echo.Context) int {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["id"].(float64)
	return int(id)
}
func GetClaimsName(c echo.Context) string {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	return name
}
