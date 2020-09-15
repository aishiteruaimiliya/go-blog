package service

import (
	"blog/model"
	"blog/model/blogs"
)

func SendMassage(message *blogs.Message) error {
	db := model.GetDB()
	defer db.Close()
	return db.Create(message).Error
}

func RecvMessage(aid string) ([]blogs.Message, error) {
	db := model.GetDB()
	defer db.Close()
	var messages []blogs.Message
	err := db.Model(&blogs.Message{}).Find(messages, "recvid=?", aid).Error
	return messages, err
}
