package service

import (
	"blog/model"
	"blog/model/blogs"
	"fmt"
)
func UpdateContact(contact *blogs.Contact)error{
	db:=model.GetDB()
	defer db.Close()
	err:=db.Model(blogs.Contact{}).Updates(contact).Error
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
