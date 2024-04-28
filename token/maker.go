package token

import (
	"fmt"
	"playground/util"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(username string) (string, error) {
	var (
		tokenString string
		err         error
	)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
	})

	if tokenString, err = token.SignedString([]byte(util.Config.EncKey)); err != nil {
		return "", err
	}
	return tokenString, err
}

func ValidateToken(tokenString string) bool {
	var (
		err   error
	)
	if _, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(util.Config.EncKey), nil
	}); err != nil {
		return false
	}
	return true
}
