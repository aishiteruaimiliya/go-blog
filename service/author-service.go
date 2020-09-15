package service

import (
	"blog/model"
	"blog/model/blogs"
)

func GetAuthorInfoByAccountId(id string) (blogs.User, error) {
	db := model.GetDB()
	defer db.Close()
	var res blogs.User
	err := db.Model(&blogs.User{}).Where(&blogs.User{Account: id}).First(&res).Error
	if err != nil {
		return blogs.User{}, err
	}
	return res, nil
}

func UpdateInfo(author *blogs.User) error {
	db := model.GetDB()
	defer db.Close()
	return db.Model(blogs.User{}).Updates(author).Error
}
