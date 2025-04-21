package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/celestix/autodeployer_api/api/common"
	"github.com/celestix/autodeployer_api/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func jwtKeyFunc(token *jwt.Token) (any, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}
	return []byte(config.Data.SecretKey), nil
}

func Auth(c *gin.Context) {
	authH := c.GetHeader("Authorization")
	if authH == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, common.ErrorResponse(common.ErrUnauthorized))
		return
	}
	authToken := strings.Split(authH, " ")
	if len(authToken) != 2 || authToken[0] != "Bearer" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "invalid token")
		return
	}
	token, err := jwt.Parse(authToken[1], jwtKeyFunc)
	if err != nil || !token.Valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, common.ErrorResponse(common.ErrInvalidToken))
		return
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, common.ErrorResponse(common.ErrInvalidToken))
		return
	}
	name := claims["name"].(string)
	c.Set("userName", name)
	pat := claims["pat"].(string)
	c.Set("gho", pat)
	c.Next()
}
