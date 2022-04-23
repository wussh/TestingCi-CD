package middleware

import (
	"kentang/key"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateToken(userID, name string) (string, error) {
	claims := jwt.MapClaims{}
	claims["userID"] = userID
	claims["name"] = name
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(key.SECRET_JWT))
}
