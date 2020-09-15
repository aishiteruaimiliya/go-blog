package service

import (
	"blog/model"
	"blog/model/blogs"
	"blog/utils"
	"fmt"
	"time"
)

// 注册
func Register(user *blogs.User) error {
	var err error
	db := model.GetDB()
	defer db.Close()

	db.Begin()
	err = db.Create(user).Error
	if err != nil {
		db.Rollback()
		return err
	}
	err = db.Create(&blogs.Contact{UserName: user.Account}).Error
	if err != nil {
		db.Rollback()
		return err
	}
	db.Commit()
	return nil
}

// 登录
func Login(user *blogs.User) (string, error) {
	db := model.GetDB()
	defer db.Close()
	var count int
	err := db.Model(&blogs.User{}).Where(user).Count(&count).Error
	fmt.Println("count is", count)
	if err != nil {
		return "", err
	}
	// todo  登陆成功后返回token，并将token放进redis
	if count >= 1 {
		generateUUID, _ := utils.GenerateTokenByAccount(user)
		err = TokenToRedis(generateUUID, user.Account, time.Second*300)
		return generateUUID, nil
	}
	return "", nil
}

// 更新密码
func UpdatePassword(username, old, new string) (bool, error) {
	db := model.GetDB()
	defer db.Close()
	auth, err := Login(&blogs.User{Account: username, Password: old})
	if err != nil {
		return false, err
	}
	if auth != "" {
		err = db.Model(&blogs.User{Account: username}).Update("password", new).Error
		if err != nil {
			return false, err
		}
	} else {
		return false, nil
	}
	return true, nil
}

// ban掉用户账号
func BannedAccount(account string) error {
	db := model.GetDB()
	defer db.Close()
	return db.Model(&blogs.User{Account: account}).Update("banned", true).Error
}
