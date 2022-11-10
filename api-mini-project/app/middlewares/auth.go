package middlewares

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var whitelist []string = make([]string, 5)

type JwtCustomClaims struct {
	ID int `json:"id"`
	jwt.StandardClaims
}

type ConfigJWT struct {
	SecretJWT       string
	ExpiresDuration int
}

func (jwtConf *ConfigJWT) Init() middleware.JWTConfig {
	return middleware.JWTConfig{
		Claims:     &JwtCustomClaims{},
		SigningKey: []byte(jwtConf.SecretJWT),
	}
}

func (jwtConf *ConfigJWT) GenerateToken(userID int) string {
	claims := JwtCustomClaims{
		userID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(int64(jwtConf.ExpiresDuration))).Unix(),
		},
	}

	//Create token with claims
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, _ := t.SignedString([]byte(jwtConf.SecretJWT))

	whitelist = append(whitelist, token)

	return token
}

func GetUser(c echo.Context) *JwtCustomClaims {
	user := c.Get("user").(*jwt.Token)

	isListed := CheckToken(user.Raw)

	if !isListed {
		return nil
	}

	claims := user.Claims.(*JwtCustomClaims)

	return claims
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

// func CreateToken(userId uint) string {
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// 		"userId": userId,
// 		"expire": time.Now().Add(time.Hour * 1).Unix(),
// 	})

// 	tokenString, err := token.SignedString([]byte("secretkey"))

// 	if err != nil {
// 		log.Fatalf("error when creating token: %v", err)
// 	}

// 	whitelist = append(whitelist, tokenString)

// 	return tokenString

// }

// func ExtractToken(e echo.Context) uint {
// 	user := e.Get("user").((*jwt.Token))

// 	isListed := CheckToken(user.Raw)
// 	if !isListed {
// 		return 0
// 	}

// 	if user.Valid {
// 		claims := user.Claims.(jwt.MapClaims)
// 		userId := claims["userId"].(float64)

// 		return uint(userId)
// 	}

// 	return 0
// }
