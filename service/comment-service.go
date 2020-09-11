package service

import (
	"blog/model"
	"blog/model/blogs"
	"fmt"
)

func AddComment(comment *blogs.Comment)error{
	db:=model.GetDB()
	defer db.Close()
	err := db.Create(comment).Error
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func DeleteComment(cid string)error{
	db:=model.GetDB()
	defer db.Close()
	err:=db.Delete(&blogs.Comment{},"cid=?",cid).Error
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func FindCommentByBlog(bid string)([]blogs.Comment,error){
	db:=model.GetDB()
	defer db.Close()
	var comments []blogs.Comment
	err:=db.Model(&blogs.Comment{}).Find(&comments,"bid=?",bid).Error
	if err != nil {
		return []blogs.Comment{},err
	}else {
		return comments,err
	}

}

func FindCommentByAuthor(aid string)([]blogs.Comment,error){
	db:=model.GetDB()
	defer db.Close()
	var comments []blogs.Comment
	err:=db.Model(&blogs.Comment{}).Find(&comments,"aid=?",aid).Error
	if err != nil {
		return []blogs.Comment{},err
	}else {
		return comments,err
	}
}
