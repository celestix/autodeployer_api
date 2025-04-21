package utils

import (
	"github.com/celestix/autodeployer_api/config"
	"github.com/golang-jwt/jwt"
)

func GenerateJWT(name, pat string) (string, error) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": name,
		"pat":  pat,
	})
	return jwtToken.SignedString([]byte(config.Data.SecretKey))
}
