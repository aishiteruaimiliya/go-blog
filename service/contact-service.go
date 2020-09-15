package service

import (
	"blog/model"
	"blog/model/blogs"
)

func UpdateContact(contact *blogs.Contact) error {
	db := model.GetDB()
	defer db.Close()
	return db.Model(blogs.Contact{}).Updates(contact).Error
}
