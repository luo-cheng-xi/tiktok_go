package jwt_utils

import (
	"github.com/golang-jwt/jwt/v5"
	"log"
	"tiktok/conf/jwt_conf"
	"time"
)

type Payload struct {
	ID uint
	jwt.RegisteredClaims
}

// GetJwt 通过用户id参数得到JWT令牌
func GetJwt(id uint) string {
	claims := Payload{
		id,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // 过期时间24小时
			IssuedAt:  jwt.NewNumericDate(time.Now()),                     // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),                     // 生效时间
		},
	}
	// 使用HS256签名算法
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	s, err := t.SignedString([]byte(jwt_conf.JwtSignedKey))
	if err != nil {
		log.Default().Println("生成jwt令牌出错")
		return ""
	}
	return s
}

// ParseJwt 解析JWT
func ParseJwt(tokenString, secretKey string) (*Payload, error) {
	t, err := jwt.ParseWithClaims(tokenString, &Payload{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}
	//检查信息是否为payload类型，如果是的话就返回其中包含的信息
	if claims, ok := t.Claims.(*Payload); ok && t.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
