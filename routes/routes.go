package routes

import (
	"blog/model/blogs"
	"blog/routes/middleware"
	"blog/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-uuid"
	"log"
	"net/http"
	"time"
)

func InitRoutes(g *gin.Engine){
	g.Use(gin.Recovery())
	g.Use(middleware.Logger)
	g.Use(middleware.Auth)
	initAccountRoute(g)
	initBlogRoute(g)
	initCommentRoute(g)
	initContactRoute(g)
	initMessageRoute(g)
	initAuthorRoute(g)
}

func initAuthorRoute(g *gin.Engine) {
	r:=g.Group("/author")
	{
		r.GET("/info", func(context *gin.Context) {
			q:=context.Query("id")
			r, err := service.GetAuthorInfoByAccountId(q)
			if err != nil {
				context.JSON(http.StatusNotFound,"未找到您的信息")
			}
			context.JSON(http.StatusOK,r)
		})
		r.PUT("/info", func(context *gin.Context) {
			author:= &blogs.Author{}
			err := context.ShouldBind(author)
			if err != nil {
				context.JSON(http.StatusMethodNotAllowed,"请仔细检查信息是否完整")
			}
			err = service.UpdateInfo(author)
			if err != nil {
				context.JSON(http.StatusServiceUnavailable,"信息修改失败")
			}
			context.JSON(http.StatusOK,author)
		})
	}
}

func initCommentRoute(g *gin.Engine) {
	comment := g.Group("/comment")
	{
		comment.POST("/comment", func(context *gin.Context) {
			c:=&blogs.Comment{}
			err := context.ShouldBind(c)
			if err != nil {
				context.JSON(http.StatusBadRequest,"请求参数有误")
			}else {
				err := service.AddComment(c)
				if err != nil {
					context.JSON(http.StatusInternalServerError,"内部错误")
				}else{
					context.JSON(http.StatusOK,c)
				}
			}
		})
		comment.DELETE("/comment", func(context *gin.Context) {
			q:=context.Query("cid")
			err := service.DeleteComment(q)
			if err != nil {
				context.JSON(http.StatusBadRequest,"请求错误")
			}else{
				context.JSON(http.StatusOK,q)
			}
		})
		comment.GET("/byBlog", func(context *gin.Context) {
			q:=context.Query("bid")
			blog, err := service.FindCommentByBlog(q)
			if err != nil {
				context.JSON(http.StatusInternalServerError,blog)
			}else {
				context.JSON(http.StatusOK,blog)
			}
		})
		comment.GET("/byAuthor", func(context *gin.Context) {
			q:=context.Query("aid")
			blog, err := service.FindCommentByAuthor(q)
			if err != nil {
				context.JSON(http.StatusInternalServerError,blog)
			}else {
				context.JSON(http.StatusOK,blog)
			}
		})
	}
}

func initMessageRoute(g *gin.Engine) {
	r:=g.Group("/message")
	{
		r.GET("/message", func(context *gin.Context) {
			// todo 此后此处需要修改aid的获取方式
			q:=context.Query("id")
			message, err := service.RecvMessage(q)
			if err != nil {
				context.JSON(http.StatusInternalServerError,"内部错误")
			}else {
				context.JSON(http.StatusOK,message)
			}
		})
		r.POST("/message", func(context *gin.Context) {
			msg:=&blogs.Message{}
			err := context.ShouldBind(msg)
			if err != nil {
				context.JSON(http.StatusBadRequest,"请求参数错误")
			}else {
				err := service.SendMassage(msg)
				if err != nil {
					context.JSON(http.StatusInternalServerError,"内部错误")
				}else {
					context.JSON(http.StatusOK,msg)
				}
			}
		})
	}
}

func initBlogRoute(g *gin.Engine) {
	blog:=g.Group("/blog")
	{
		blog.GET("/list", func(context *gin.Context) {
			param := context.Query("id")
			myBlog, err := service.GetMyBlog(param)
			if err != nil {
				context.JSON(http.StatusBadRequest,"请求失败")
			}else{
				context.JSON(http.StatusNotFound, myBlog)
			}
		})
		blog.GET("/detail", func(context *gin.Context) {
			para:=context.Query("id")
			b, err := service.GetBlogById(para)
			if err != nil {
				context.JSON(http.StatusNotFound,nil)
			}else{
				context.JSON(http.StatusOK,b)
			}

		})
		blog.GET("/all", func(context *gin.Context) {
			q:=context.Query("id")
			allBlogs, err := service.GetAllBlogs(q)
			if err != nil {
				context.JSON(http.StatusInternalServerError,"内部错误")
			}else {
				context.JSON(http.StatusOK,allBlogs)
			}
		})
		blog.POST("/blog", func(context *gin.Context) {
			blog:=&blogs.Blog{}
			err := context.ShouldBind(blog)
			blog.BID, _ = uuid.GenerateUUID()
			blog.TS = time.Now()
			if err != nil {
				context.JSON(http.StatusBadRequest,"请求参数不正确")
			}else{
				err = service.SendBlog(blog)
				if err != nil {
					context.JSON(http.StatusInternalServerError,"内部错误")
				}else {
					context.JSON(http.StatusOK,blog)
				}
			}

		})
		blog.PUT("/blog", func(context *gin.Context) {
			blog:=&blogs.Blog{}
			err := context.ShouldBind(blog)
			blog.TS = time.Now()
			if err != nil {
				context.JSON(http.StatusBadRequest,"请求参数不正确")
			}else{
				err = service.UpdateBlog(blog)
				if err != nil {
					context.JSON(http.StatusInternalServerError,"内部错误")
				}else {
					context.JSON(http.StatusOK,blog)
				}
			}
		})
		blog.DELETE("/blog", func(context *gin.Context) {
			r := context.Query("bid")
			err := service.DeleteBlog(&blogs.Blog{BID: r})
			if err != nil {
				context.JSON(http.StatusInternalServerError,"内部错误")
			}else {
				context.JSON(http.StatusOK,r)
			}
		})
	}
}

func initAccountRoute(g *gin.Engine) {
	user:=g.Group("/account")
	{
		user.POST("/register", func(context *gin.Context) {
			user := &blogs.Account{}
			err := context.ShouldBind(user)
			if err != nil {
				fmt.Println(err)
				context.JSON(401,"注册信息不完整")
			}
			err = service.Register(user)
			if err != nil {
				context.JSON(401,"注册失败")
			}else {
				context.JSON(http.StatusOK,"注册成功")
			}

		})
		user.POST("/login", func(context *gin.Context) {
			user:=&blogs.Account{}
			err:=context.ShouldBind(user)
			if err != nil {
				log.Print(err)
				context.JSON(401,"登陆失败")
			}
			fmt.Println(user)
			r, err := service.Login(user)
			fmt.Println(user)
			if err != nil {
				log.Println(err)
				context.JSON(401,"登陆失败")
			}
			if r{
				context.JSON(http.StatusOK,"登陆成功")
			}else {
				context.JSON(http.StatusUnauthorized,"登陆失败")
			}
		})
		g.GET("/hello", func(context *gin.Context) {
			context.String(http.StatusOK,"%s","hello")
		})
	}
}

func initContactRoute(g *gin.Engine) {
	r:=g.Group("/contact")
	{
		r.PUT("/contact", func(context *gin.Context) {
			contact:=&blogs.Contact{}
			err := context.ShouldBind(contact)
			if err != nil {
				context.JSON(http.StatusBadRequest,"请求参数不正确")
			}else{
				err := service.UpdateContact(contact)
				if err != nil {
					context.JSON(http.StatusInternalServerError,"内部错误")
				}else {
					context.JSON(http.StatusOK,contact)
				}
			}
		})
	}
}