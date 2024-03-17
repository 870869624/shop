package models

import (
	"errors"
	"shop/db"
	"time"

	"gorm.io/gorm"
)

type UserKind int

const (
	//普通用户
	GeneralUser UserKind = 1
	//管理员
	Administrator UserKind = 0
)

type User struct {
	ID          int
	Username    string
	Password    string
	Phone       string
	Gender      int8
	Age         uint
	RefreshedAt time.Time
	UserKind
}

// 检验用户是否已经存在
func (u *User) CheckSameExists() bool {
	var count int64
	db := db.Connect()
	db.Table("users").Where("username = ?", u.Username).Count(&count)
	return count > 0
}

// 创建用户
func (u *User) Creat() error {
	db := db.Connect()
	if u.CheckSameExists() {
		return errors.New("已存在该用户")
	}
	result := db.Create(&u)
	if result.Error != nil {
		return errors.New("无法创建")
	}
	return nil
}

func (u *User) Delete() {
	db := db.Connect()
	db.Where("username = ?", u.Username).Delete(&u)
}
func (u *User) Updata(str string) error {
	var users User
	db := db.Connect()
	result := db.Where(map[string]interface{}{"username": str}).First(&users)
	//理论上token里包含用户信息!!!!
	if result.Error != nil {
		return errors.New("该用户不存在")
	}
	result.Updates(map[string]interface{}{"password": u.Password, "phone": u.Phone})
	err := u.BeforeUpdate(result)
	if err != nil {
		return err
	}
	return nil
}

// 检查是否修改过信息
func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	// // if Role changed
	// if tx.Statement.Changed("user6name") {
	// 	return errors.New("role not allowed to change")
	// }

	// if tx.Statement.Changed("password", "phone") { // if Name or Role changed
	// 	tx.Statement.SetColumn("Age", 18)
	// }

	// if any fields changed
	if tx.Statement.Changed() {
		tx.Statement.SetColumn("refreshedat", time.Now())
		return nil
	}
	return errors.New("信息未修改")
}
