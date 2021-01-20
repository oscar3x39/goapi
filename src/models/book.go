package models

import (
	"goapi/config"
)

type Book struct {
	Name 	 string `json:"name"`
	Author 	 string `json:"author"`
	Category string `json:"category"`
}

func GetAllBook(b *[]Book) (err error) {
	if err = config.DB.Find(b).Error; err != nil {
		return err
	}
	return nil
}