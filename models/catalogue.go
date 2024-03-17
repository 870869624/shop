package models

import "shop/db"

type Catalogue struct {
	ID            int
	Cataloguename string
	Icon          string
	Order         int
}

func (c *Catalogue) DoesSameExists() bool {
	var count int64
	db := db.Connect()
	db.Table("catalogue").Where("cataloguename = ?", c.Cataloguename).Count(&count)
	return count > 0
}
func (c *Catalogue) CreateCatalogue() bool {
	return true
}
