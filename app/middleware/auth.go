package middleware

import (
	"final_project/controllers"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type JwtMyClaims struct {
	ID int `json:"id"`
	jwt.StandardClaims
}

type ConfigJWT struct {
	SecretJWT       string
	ExpiresDuration int
}

func (jwtConf *ConfigJWT) Init() middleware.JWTConfig {
	return middleware.JWTConfig{
		Claims:     &JwtMyClaims{},
		SigningKey: []byte(jwtConf.SecretJWT),
		ErrorHandlerWithContext: middleware.JWTErrorHandlerWithContext(func(e error, c echo.Context) error {
			return controllers.NewErrorResponse(c, http.StatusForbidden, e)
		}),
	}
}

func (jwtConf *ConfigJWT) GenerateToken(ID int) (string, error) {
	claims := JwtMyClaims{
		ID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(int64(jwtConf.ExpiresDuration))).Unix(),
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := t.SignedString([]byte(jwtConf.SecretJWT))

	return token, err
}
func GetClaimsUserId(c echo.Context) int {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtMyClaims)
	id := claims.ID
	return int(id)
}
