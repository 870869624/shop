package entities

//检验值是否符合要求
import (
	"errors"
	"time"
)

type UserLoginRequestPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserLoginResponPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type UserRegisterPayload struct {
	ID          int       `json:"id"`
	Username    string    `json:"username"`
	Password    string    `json:"password"`
	Phone       string    `json:"phone"`
	Gender      int8      `json:"gender"`
	Age         uint      `json:"age"`
	RefreshedAt time.Time `json:"refreshedat"`
}

func (u *UserRegisterPayload) Validate() error {
	if u.Username == "" {
		return errors.New("用户名长度有误")
	}
	if len(u.Phone) != 11 {
		return errors.New("手机号长度有误")
	}
	if len(u.Password) < 8 || len(u.Password) > 16 {
		return errors.New("密码长度有误")
	}
	return nil
}

func (u *UserLoginRequestPayload) Validate() error {
	if u.Username == "" {
		return errors.New("用户名长度有误")
	}
	if len(u.Password) < 8 || len(u.Password) > 16 {
		return errors.New("密码长度有误")
	}
	return nil
}
