package middleware

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type ConfigJWT struct {
	SecretKey   string
	ExpiredTime int
}

type JwtCustomClaims struct {
	UserId uint `json:"id"`
	jwt.RegisteredClaims
}



func (jwtConf *ConfigJWT) Init() echojwt.Config {
	return echojwt.Config{
		NewClaimsFunc:     func(c echo.Context) jwt.Claims {
			return new(JwtCustomClaims)
		},
		SigningKey: []byte(jwtConf.SecretKey),
		
	}
}

func (configJwt *ConfigJWT) GenererateToken(userId uint) (string) {
	claims := JwtCustomClaims{
		userId,
		jwt.RegisteredClaims{
			ExpiresAt:jwt.NewNumericDate(time.Now().Local().Add(time.Hour * time.Duration(configJwt.ExpiredTime))),
		},
	}

	// Create token with claims
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, _ := t.SignedString([]byte(configJwt.SecretKey))
	return token
}