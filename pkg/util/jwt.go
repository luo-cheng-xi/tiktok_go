package util

import (
	"github.com/golang-jwt/jwt/v5"
	"tiktok/internal/conf"
	"time"
)

type JwtUtil struct {
	jwtConf *conf.JwtConfig
}

func GetJwtUtil(jc *conf.JwtConfig) *JwtUtil {
	return &JwtUtil{
		jwtConf: jc,
	}
}

type Payload struct {
	ID uint
	jwt.RegisteredClaims
}

// GetJwt 通过用户id参数得到JWT令牌
func (rx *JwtUtil) GetJwt(id uint) string {
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
	s, err := t.SignedString([]byte(rx.jwtConf.SignedKey))
	if err != nil {
		return ""
	}
	return s
}

// ParseJwt 解析JWT
func (rx *JwtUtil) ParseJwt(tokenString string) (*Payload, error) {
	t, err := jwt.ParseWithClaims(tokenString, &Payload{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(rx.jwtConf.SignedKey), nil
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
