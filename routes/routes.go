package routes

import (
	"blog/model/blogs"
	"blog/routes/middleware"
	"blog/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-uuid"
	"net/http"
	"strings"
	"time"
)

func InitRoutes(g *gin.Engine) {
	g.Use(gin.Recovery())
	g.Use(middleware.Logger)
	g.Use(middleware.Auth)
	g.Use(middleware.Secure)
	g.Use(middleware.Options)
	initAccountRoute(g)
	initBlogRoute(g)
	initCommentRoute(g)
	initContactRoute(g)
	initMessageRoute(g)
}

func initCommentRoute(g *gin.Engine) {
	comment := g.Group("/comment")
	{
		comment.POST("/comment", func(context *gin.Context) {
			if !validateRequest(context) {
				return
			}
			c := &blogs.Comment{}
			err := context.ShouldBind(c)
			if err != nil {
				context.JSON(http.StatusBadRequest, gin.H{"msg": "请求参数有误"})
			} else {
				err := service.AddComment(c)
				if err != nil {
					context.JSON(http.StatusInternalServerError, gin.H{"msg": "服务器内部错误"})
				} else {
					context.JSON(http.StatusOK, c)
				}
			}
		})
		comment.DELETE("/comment", func(context *gin.Context) {
			if !validateRequest(context) {
				return
			}
			q := context.Query("cid")
			err := service.DeleteComment(q)
			if err != nil {
				context.JSON(http.StatusBadRequest, gin.H{"msg": "请求错误"})
			} else {
				context.JSON(http.StatusOK, q)
			}
		})
		comment.GET("/byBlog", func(context *gin.Context) {
			if !validateRequest(context) {
				return
			}
			q := context.Query("bid")
			blog, err := service.FindCommentByBlog(q)
			if err != nil {
				context.JSON(http.StatusInternalServerError, blog)
			} else {
				context.JSON(http.StatusOK, blog)
			}
		})
		comment.GET("/byAuthor", func(context *gin.Context) {
			if !validateRequest(context) {
				return
			}
			q := context.Query("aid")
			blog, err := service.FindCommentByAuthor(q)
			if err != nil {
				context.JSON(http.StatusInternalServerError, blog)
			} else {
				context.JSON(http.StatusOK, blog)
			}
		})
	}
}

func initMessageRoute(g *gin.Engine) {
	r := g.Group("/message")
	{
		r.GET("/message", func(context *gin.Context) {
			if !validateRequest(context) {
				return
			}
			// todo 此后此处需要修改aid的获取方式
			q := context.GetString("aid")
			message, err := service.RecvMessage(q)
			if err != nil {
				context.JSON(http.StatusInternalServerError, gin.H{"msg": "内部错误"})
			} else {
				context.JSON(http.StatusOK, message)
			}
		})
		r.POST("/message", func(context *gin.Context) {
			if !validateRequest(context) {
				return
			}
			msg := &blogs.Message{}
			err := context.ShouldBind(msg)
			if err != nil {
				context.JSON(http.StatusBadRequest, gin.H{"msg": "请求参数错误"})
			} else {
				err := service.SendMassage(msg)
				if err != nil {
					context.JSON(http.StatusInternalServerError, gin.H{"msg": "内部错误"})
				} else {
					context.JSON(http.StatusOK, msg)
				}
			}
		})
	}
}

func initBlogRoute(g *gin.Engine) {
	blog := g.Group("/blog")
	{
		blog.GET("/list", func(context *gin.Context) {
			if !validateRequest(context) {
				return
			}
			param := context.Query("id")
			myBlog, err := service.GetMyBlog(param)
			if err != nil {
				context.JSON(http.StatusBadRequest, gin.H{"msg": "请求失败"})
			} else {
				context.JSON(http.StatusNotFound, myBlog)
			}
		})
		blog.GET("/detail", func(context *gin.Context) {
			if !validateRequest(context) {
				return
			}
			para := context.Query("id")
			b, err := service.GetBlogById(para)
			if err != nil {
				context.JSON(http.StatusNotFound, nil)
			} else {
				context.JSON(http.StatusOK, b)
			}

		})
		blog.GET("/all", func(context *gin.Context) {
			if !validateRequest(context) {
				return
			}
			q := context.Query("id")
			allBlogs, err := service.GetAllBlogs(q)
			if err != nil {
				context.JSON(http.StatusInternalServerError, gin.H{"msg": "内部错误"})
			} else {
				context.JSON(http.StatusOK, allBlogs)
			}
		})
		blog.POST("/blog", func(context *gin.Context) {
			if !validateRequest(context) {
				return
			}
			blog := &blogs.Blog{}
			err := context.ShouldBind(blog)
			blog.BID, _ = uuid.GenerateUUID()
			blog.TS = time.Now()
			if err != nil {
				context.JSON(http.StatusBadRequest, gin.H{"msg": "请求参数不正确"})
			} else {
				err = service.SendBlog(blog)
				if err != nil {
					context.JSON(http.StatusInternalServerError, gin.H{"msg": "内部错误"})
				} else {
					context.JSON(http.StatusOK, blog)
				}
			}

		})
		blog.PUT("/blog", func(context *gin.Context) {
			if !validateRequest(context) {
				return
			}
			blog := &blogs.Blog{}
			err := context.ShouldBind(blog)
			blog.TS = time.Now()
			if err != nil {
				context.JSON(http.StatusBadRequest, gin.H{"msg": "请求参数不正确"})
			} else {
				err = service.UpdateBlog(blog)
				if err != nil {
					context.JSON(http.StatusInternalServerError, gin.H{"msg": "内部错误"})
				} else {
					context.JSON(http.StatusOK, blog)
				}
			}
		})
		blog.GET("/view_count", func(context *gin.Context) {
			if !validateRequest(context) {
				return
			}
			bid := context.Query("bid")
			if bid == "" {
				context.JSON(http.StatusNotFound, gin.H{"msg": "文章不存在"})
				return
			} else {
				times, err := service.GetViewTimes(bid)
				if err != nil {
					fmt.Println(err)
					context.JSON(http.StatusInternalServerError, gin.H{"msg": "内部错误"})
				} else {
					context.JSON(http.StatusOK, gin.H{"count": times})
				}
			}
		})
		blog.PUT("/view_count", func(context *gin.Context) {
			if !validateRequest(context) {
				return
			}
			bid := context.Query("bid")
			if bid == "" {
				context.JSON(http.StatusNotFound, gin.H{"msg": "文章不存在"})
				return
			} else {
				times, err := service.AddViewTimes(bid)
				if err != nil {
					fmt.Println(err)
					context.JSON(http.StatusInternalServerError, gin.H{"msg": "内部错误"})
				} else {
					context.JSON(http.StatusOK, gin.H{"count": times})
				}
			}
		})
		blog.DELETE("/blog", func(context *gin.Context) {
			if !validateRequest(context) {
				return
			}
			r := context.Query("bid")
			err := service.DeleteBlog(&blogs.Blog{BID: r, Author: context.GetString("aid")})
			if err != nil {
				context.JSON(http.StatusInternalServerError, gin.H{"msg": "内部错误"})
			} else {
				context.JSON(http.StatusOK, r)
			}
		})
	}
}

func initAccountRoute(g *gin.Engine) {
	user := g.Group("/account")
	{
		user.POST("/register", func(context *gin.Context) {
			user := &blogs.User{}
			err := context.ShouldBind(user)
			if err != nil {
				fmt.Println(err)
				context.JSON(401, gin.H{"msg": "注册信息不完整"})
			}
			err = service.Register(user)
			if err != nil {
				if strings.Index(err.Error(), "违反唯一") >= 0 {
					context.JSON(http.StatusOK, gin.H{"msg": "该用户名已被注册"})
					return
				}
				context.JSON(401, gin.H{"msg": "注册失败"})
				return
			} else {
				context.JSON(http.StatusOK, gin.H{"msg": "注册成功"})
				return
			}

		})
		user.POST("/login", func(context *gin.Context) {
			aid := context.GetString("aid")
			if aid != "" {
				context.Header("token", context.GetString("token"))
				context.JSON(http.StatusOK, gin.H{"msg": "登陆成功"})
				return
			}
			user := &blogs.User{}
			// 坑点，bind不会返回错误
			err := context.ShouldBind(user)
			fmt.Println(user)
			if err != nil || user.Password == "" || user.Account == "" {
				// log.Print(err)
				context.JSON(401, gin.H{"msg": "登陆失败"})
				return
			}
			fmt.Println(user)
			r, err := service.Login(user)
			fmt.Println("err is", err)
			if err != nil {
				// log.Println(err)
				context.JSON(401, gin.H{"msg": "登陆失败"})
				return
			}
			if r != "" {
				context.Header("token", r)
				context.JSON(http.StatusOK, gin.H{"msg": "登陆成功"})
				return
			} else {
				context.JSON(http.StatusUnauthorized, gin.H{"msg": "登陆失败"})
				return
			}
		})
		user.PUT("/info", func(context *gin.Context) {
			if !validateRequest(context) {
				return
			}
			type r struct {
				Account string `json:"account"`
				OldPass string `json:"old_pass"`
				NewPass string `json:"new_pass"`
			}
			blog := &r{}
			err := context.ShouldBind(blog)
			if err != nil {
				context.JSON(http.StatusBadRequest, gin.H{"msg": "请求参数错误"})
			} else {
				res, err := service.UpdatePassword(blog.Account, blog.OldPass, blog.NewPass)
				if err != nil {
					context.JSON(http.StatusInternalServerError, gin.H{"msg": "内部错误"})
				} else if res {
					context.JSON(http.StatusOK, blog)
				} else {
					context.JSON(http.StatusMethodNotAllowed, gin.H{"msg": "密码不正确"})
				}
			}
		})
	}
}

func initContactRoute(g *gin.Engine) {
	r := g.Group("/contact")
	{
		r.PUT("/contact", func(context *gin.Context) {
			if !validateRequest(context) {
				return
			}
			contact := &blogs.Contact{}
			err := context.ShouldBind(contact)
			if err != nil {
				context.JSON(http.StatusBadRequest, gin.H{"msg": "请求参数不正确"})
			} else {
				err := service.UpdateContact(contact)
				if err != nil {
					context.JSON(http.StatusInternalServerError, gin.H{"msg": "内部错误"})
				} else {
					context.JSON(http.StatusOK, contact)
				}
			}
		})
	}
}

func validateRequest(context *gin.Context) bool {
	aid := context.GetString("aid")
	//fmt.Println(aid,"in validateRequest")
	if aid == "" {
		context.JSON(http.StatusUnauthorized, gin.H{"msg": "未登录，请登录"})
		return false
	}
	return true
}
