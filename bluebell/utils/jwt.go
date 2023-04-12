package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// 过期时间

const TokenExpireDuration = time.Hour * 2

// 盐

const MySecret = "bluebell12131"

type MyClaims struct {
	UserId   int64  `json:"userId"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// 生成token

func GenToken(userId int64, username string) (string, error) {
	c := MyClaims{
		UserId:   userId,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    MySecret,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(MySecret)
}

// 解析token

func ParseToken(tokenSting string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenSting, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})

	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, err
	}
	return nil, errors.New("token解析失败")
}
