package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"tiktok/setting"
)

// GetJwt 通过claims参数得到JWT令牌
func GetJwt(claims jwt.Claims) (string, error) {
	key := setting.JwtSignedKey
	s, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(key)
	if err != nil {
		return "", errors.New("claims invalid")
	}
	return s, nil
}

// ParseJwt 通过解析JWT令牌得到claims
func ParseJwt(str string, claims jwt.Claims) (jwt.Claims, error) {
	token, err := jwt.ParseWithClaims(str, claims, func(token *jwt.Token) (interface{}, error) {
		return setting.JwtSignedKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("claim invalid")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid claim type")
	}
	return claims, nil
}
