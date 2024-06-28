package config

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtSecret = []byte("2024xiaoxueqi") // 这里可以改为你自己的密钥

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenerateToken 生成JWT Token
func GenerateToken(username string) (string, error) {
	now := time.Now()
	expirationTime := now.Add(24 * time.Hour)
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  now.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// ParseToken 解析JWT Token
func ParseToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, err
	}
	return claims, nil
}
