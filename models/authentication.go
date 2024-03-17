package models

import (
	"errors"
	"shop/db"
)

type Authentication struct {
	Username string
	Password string
}
type users struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Gender   int8   `json:"gender"`
	Age      uint   `age:"age"`
}

// 登录信息拿到加密密码然后对比数据库然后生成token
func (a *Authentication) Signin() (string, error) {
	var users users
	db := db.Connect()
	password := Encrypt(a.Password)
	request := db.Where(map[string]interface{}{"username": a.Username, "password": password}).First(&users)
	if request.Error != nil {
		err := db.Where(map[string]interface{}{"username": a.Username}).First(&users)
		if err.Error != nil {
			return "", errors.New("该用户不存在")
		}
		return "", errors.New("密码错误")
	}
	return a.CreateJwt()
}
func (a *Authentication) CreateJwt() (string, error) {
	return GenToken(a)
}
