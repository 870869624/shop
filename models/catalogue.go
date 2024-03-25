package models

import (
	"errors"
	"shop/db"
	"time"
)

type Catalogue struct {
	ID            int        `json:"id"`
	Cataloguename string     `json:"cataloguename"`
	Icon          string     `json:"icon"`
	Order         int        `json:"order"`
	CreatTime     *time.Time `json:"creat_at"`
}

func (c *Catalogue) DoesSameExists() bool {
	var count int64
	db := db.Connect()
	db.Table("catalogue").Where("cataloguename = ?", c.Cataloguename).Count(&count)
	return count > 0
}
func (c *Catalogue) Validata() error {
	if c.Cataloguename == "" {
		return errors.New("类目信息错误")
	}
	if c.Icon == "" {
		return errors.New("图标信息错误")
	}
	//排序信息，判断是否为空
	if c.Order == 0 {
		return errors.New("顺序信息错误")
	}
	return nil
}
func (c *Catalogue) CreateCatalogue() error {
	db := db.Connect()
	if !c.DoesSameExists() {
		if err := c.Validata(); err != nil {
			return err
		}
		if err := db.Create(&c); err != nil {
			return err.Error
		}
	}
	return errors.New("已经存在相同类目")
}
