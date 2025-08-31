package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var JWTSECRET = "sfsdfsdfjsdfljiofjwefnljfsfdlmfdlfkslpkjfjsdhfoisdfj"

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(JWTSECRET))
}

func VerifyToken(token string) (int64, error) {
	passToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Unexpected signing method")
		}
		return []byte(JWTSECRET), nil
	})

	if err != nil {
		return 0, errors.New("Could not parse token")
	}

	tokenIsValid := passToken.Valid

	if !tokenIsValid {
		return 0, errors.New("Invalid Token")
	}

	claim, ok := passToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("Invalid Token Claim")
	}

	email := claim["email"].(string)
	//userId := int64(claim["userId"].(float64))
	fmt.Println(email)
	userId := int64(claim["userId"].(float64))
	return userId, nil
}
