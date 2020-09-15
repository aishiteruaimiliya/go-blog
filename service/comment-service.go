package service

import (
	"blog/model"
	"blog/model/blogs"
)

func AddComment(comment *blogs.Comment) error {
	db := model.GetDB()
	defer db.Close()
	return db.Create(comment).Error
}

func DeleteComment(cid string) error {
	db := model.GetDB()
	defer db.Close()
	return db.Delete(&blogs.Comment{}, "cid=?", cid).Error
}

func FindCommentByBlog(bid string) ([]blogs.Comment, error) {
	db := model.GetDB()
	defer db.Close()
	var comments []blogs.Comment
	err := db.Model(&blogs.Comment{}).Find(&comments, "bid=?", bid).Error
	return comments, err

}

func FindCommentByAuthor(aid string) ([]blogs.Comment, error) {
	db := model.GetDB()
	defer db.Close()
	var comments []blogs.Comment
	err := db.Model(&blogs.Comment{}).Find(&comments, "aid=?", aid).Error
	return comments, err
}
