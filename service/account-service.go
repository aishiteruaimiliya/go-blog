package service

import (
	"blog/model"
	"blog/model/blogs"
	"fmt"
	"github.com/hashicorp/go-uuid"
	"time"
)
// 注册
func Register(author *blogs.Account)error{
	var err error
	author.AID ,err= uuid.GenerateUUID()
	if err!=nil{
		fmt.Println(err)
	}
	db:=model.GetDB()
	defer db.Close()

	db.Begin()
	err = db.Create(author).Error
	if err!=nil{
		db.Rollback()
		return err
	}
	err = db.Create(&blogs.Author{Aid: author.AID}).Error
	if err!=nil{
		db.Rollback()
		return err
	}
	err = db.Create(&blogs.Contact{AID: author.AID}).Error
	if err!=nil{
		db.Rollback()
		return err
	}
	db.Commit()
	return nil
}
// 登录
func Login(user * blogs.Account)(bool,error) {
	db:=model.GetDB()
	defer db.Close()
	var count int
	err := db.Model(&blogs.Account{}).Where(user).Count(&count).Error
	generateUUID, _ := uuid.GenerateUUID()
	err = TokenToRedis(generateUUID, user.Account, time.Second*300)
	if err != nil {
		return false,err
	}
	fmt.Println(generateUUID)
	// todo  登陆成功后返回token，并将token放进redis
	if count>=1{
		return true,nil
	}
	return false,nil
}

// 更新密码
func UpdatePassword(username,old,new string)(bool,error){
	db:=model.GetDB()
	defer db.Close()
	auth, err := Login(&blogs.Account{AID: username,Password: old})
	if err != nil {
		return false,err
	}
	if auth{
		err = db.Model(&blogs.Account{Account: username}).Update("password", new).Error
		if err!=nil{
			return false,err
		}
	}else {
		return false,nil
	}
	return true,nil
}
// ban掉用户账号
func BannedAccount(account string)(bool,error){
	db:=model.GetDB()
	defer db.Close()
	err := db.Model(&blogs.Account{Account: account}).Update("banned", true).Error
	if err != nil {
		return false,err
	}
	return true,nil
}
// 通过账号查询用户信息
func FindAuthorByAccount(aid string) blogs.Author {
	db:=model.GetDB()
	defer db.Close()
	b:=blogs.Author{}
	err:=db.Model(&blogs.Author{}).Where(&blogs.Author{Aid: aid}).First(&b).Error
	if err != nil {
		return blogs.Author{}
	}else {
		return b
	}
}