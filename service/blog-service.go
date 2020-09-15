package service

import (
	"blog/model"
	"blog/model/blogs"
)

func GetAllBlogs(author string) ([]blogs.Blog, error) {
	db := model.GetDB()
	defer db.Close()
	var res []blogs.Blog
	err := db.Find(&res, "author=?", author).Error
	return res, err
}
func GetMyBlog(aid string) ([]blogs.Blog, error) {
	db := model.GetDB()
	defer db.Close()
	var res []blogs.Blog
	err := db.Find(&res, "aid=?", aid).Error
	return res, err
}
func SendBlog(blog *blogs.Blog) error {
	db := model.GetDB()
	defer db.Close()
	err := db.Create(blog).Error
	return err
}

func DeleteBlog(blog *blogs.Blog) error {
	db := model.GetDB()
	defer db.Close()
	err := db.Delete(&blogs.Blog{}).Where(blog).Error
	return err
}

func UpdateBlog(blog *blogs.Blog) error {
	db := model.GetDB()
	defer db.Close()
	err := db.Model(&blogs.Blog{}).Updates(blog).Error
	return err
}

func GetBlogById(bid string) (blogs.Blog, error) {
	db := model.GetDB()
	defer db.Close()
	blog := blogs.Blog{BID: bid}
	err := db.Model(&blogs.Blog{}).Find(&blog).Where("bid=?", bid).Error
	return blog, err
}
