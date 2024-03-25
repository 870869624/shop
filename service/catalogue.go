package service

import (
	"shop/models"
)

func CreateCatalogue(c *models.Catalogue) error {
	if err := c.CreateCatalogue(); err != nil {
		return err
	}
	return nil
}
