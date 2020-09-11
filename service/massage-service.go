package service

import (
	"blog/model"
	"blog/model/blogs"
	"fmt"
)

func SendMassage(message *blogs.Message)error{
	db := model.GetDB()
	defer db.Close()
	err := db.Create(message).Error
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func RecvMessage(aid string)([]blogs.Message,error){
	db := model.GetDB()
	defer db.Close()
	var messages []blogs.Message
	err:=db.Model(&blogs.Message{}).Find(messages,"recvid=?",aid).Error
	if err != nil {
		fmt.Println(err)
		return []blogs.Message{},err
	}
	return messages,nil
}
