package models

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// 签证申明，username为自定义,还没有判定用户等级
type UserAuthclaim struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.RegisteredClaims
}

var key = []byte("独一无二")

const TokenExpireDuration = 2 * time.Hour

// 生成token
func GenToken(u *Authentication) (string, error) {
	claim := UserAuthclaim{
		Username: u.Username,
		Password: u.Password,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpireDuration)),
			Issuer:    "景海军",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tk, err := token.SignedString(key)
	if err != nil {
		return "", err
	}
	return tk, nil
}

// 解析token,返回用户信息，拿去数据库做对比
func ParseToken(tokenString string) (*UserAuthclaim, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserAuthclaim{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil {
		return nil, err
	}
	claim, ok := token.Claims.(*UserAuthclaim)
	if !ok {
		return nil, errors.New("未知的的claim格式")
	}
	return claim, nil
}
