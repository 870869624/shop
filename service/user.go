package service

import (
	"errors"
	"shop/entities"
	"shop/models"
)

// 创建用户
func CreateUser(u entities.UserRegisterPayload) error {
	res := models.Encrypt(u.Password)
	user := models.User{
		Username: u.Username,
		Password: string(res),
		Phone:    u.Phone,
		Gender:   u.Gender,
		Age:      u.Age,
		Email:    u.Email,
		UserKind: u.UserKind,
	}
	err1 := user.Creat()
	if err1 != nil {
		return err1
	}
	return nil

}

// 删除用户
func DeleteUser(u entities.UserLoginRequestPayload) error {
	user := models.User{
		Username: u.Username,
		//后期应该改成加密密码
		Password: string(u.Password),
	}
	//如果为ok说明该用户还存在，可以删除
	if !user.CheckSameExists() {
		return errors.New("注销失败，该用户不存在")
	}
	user.Delete()
	return nil
}

// 更新用户信息//用户名不应该被更新，被更新的是昵称
func UpdataUser(u entities.UserRegisterPayload, str string) error {
	if u.Username != str {
		return errors.New("信息错误，用户名不匹配")
	}
	var user = models.User{
		Username: u.Username,
		Password: models.Encrypt(u.Password),
		Phone:    u.Phone,
		Gender:   u.Gender,
		Age:      u.Age,
	}
	if err := user.Updata(str); err != nil {
		return err
	}
	return nil
}
