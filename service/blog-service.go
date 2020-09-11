package service

import (
	"blog/model"
	"blog/model/blogs"
	"fmt"
)

func GetAllBlogs(author string)([]blogs.Blog,error){
	db:=model.GetDB()
	defer db.Close()
	var res []blogs.Blog
	err := db.Find(&res,"author=?",author).Error
	if err != nil {
		fmt.Println(err)
		return []blogs.Blog{},err
	}
	return res,nil
}
func GetMyBlog(aid string)([]blogs.Blog,error){
	db:=model.GetDB()
	defer db.Close()
	var res []blogs.Blog
	err:=db.Find(&res,"aid=?",aid).Error
	if err != nil {
		fmt.Println(err)
		return []blogs.Blog{},err
	}
	return res,nil
}
func SendBlog(blog *blogs.Blog)error{
	db:=model.GetDB()
	defer db.Close()
	err := db.Create(blog).Error
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func DeleteBlog(blog *blogs.Blog)error{
	db:=model.GetDB()
	defer db.Close()
	err:=db.Delete(&blogs.Blog{},"bid=?",blog.BID).Error
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func UpdateBlog(blog *blogs.Blog)error{
	db:=model.GetDB()
	defer db.Close()
	err:=db.Model(&blogs.Blog{}).Updates(blog).Error
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func GetBlogById(bid string)(blogs.Blog,error){
	db:=model.GetDB()
	defer db.Close()
	blog:=blogs.Blog{BID: bid}
	err:=db.Model(&blogs.Blog{}).Find(&blog).Where("bid=?",bid).Error
	if err != nil {
		fmt.Println(err)
		return blogs.Blog{},err
	}
	return blog,nil
}