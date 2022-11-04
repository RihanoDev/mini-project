package auth

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

var whitelist []string = make([]string, 5)

func CreateToken(userId uint) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"expire": time.Now().Add(time.Hour * 1).Unix(),
	})

	tokenString, err := token.SignedString([]byte("secretkey"))

	if err != nil {
		log.Fatalf("error when creating token: %v", err)
	}

	whitelist = append(whitelist, tokenString)

	return tokenString

}

func ExtractToken(e echo.Context) uint {
	user := e.Get("user").((*jwt.Token))

	isListed := CheckToken(user.Raw)
	if !isListed {
		return 0
	}

	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["userId"].(float64)

		return uint(userId)
	}

	return 0
}

func CheckToken(token string) bool {
	for _, tkn := range whitelist {
		if tkn == token {
			return true
		}
	}

	return false
}

func Logout(token string) bool {
	for idx, tkn := range whitelist {
		if tkn == token {
			whitelist = append(whitelist[:idx], whitelist[idx+1:]...)
		}
	}

	return true
}
