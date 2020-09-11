package service

import (
	"blog/model"
	"blog/model/blogs"
)

func GetAuthorInfoByAccountId(id string)(blogs.Author,error){
	db:=model.GetDB()
	defer db.Close()
	var res blogs.Author
	err := db.Model(&blogs.Author{}).Where(&blogs.Author{Aid: id}).First(&res).Error
	if err != nil {
		return blogs.Author{},err
	}
	return res,nil
}

func UpdateInfo(author *blogs.Author)error{
	db:=model.GetDB()
	defer db.Close()
	return db.Model(blogs.Author{}).Updates(author).Error
}

