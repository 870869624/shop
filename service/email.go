package service

import (
	"context"
	"shop/conf"
	"shop/db"
	"shop/models"
	"strings"
)

type SendEmailService struct {
	Email         string `json:"email"`
	Password      string `json::"password"`
	OperationType uint   `json:"operationtype"`
	//1.绑定邮箱 2.解绑 3.修改密码
}

func (s *SendEmailService) Send(ctx context.Context, username string) error {
	db := db.Connect()
	// code := e.success
	var (
		address string
		notice  model.Notice
	)
	token, err := models.EmailToken(username, s.OperationType, s.Email, s.Password)
	if err != nil {
		return err
	}
	notice, err := db.Model(&model.Notice{}).Where("username = ?", username).First(&notice)
	address = conf.ValidEmail + token //发送方
	mailStr := notice.Text
	mialTex := strings.Replace(mailStr, "Email", address, -1)
	m := mail.NewMessage()
	m.SetHeader("Form", conf.SmtpEmail)
	return nil
}
