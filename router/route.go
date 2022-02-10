package router

import (
	"blog1222-go/api/articles"
	"blog1222-go/api/uploads"
	"blog1222-go/api/users"
	"blog1222-go/config"
	"net/http"

	"github.com/gin-gonic/gin"

	"blog1222-go/middleware"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
)

type Router struct {
	port string
}

func (r *Router) Init() {
	router := gin.Default()
	router.StaticFS("/images", http.Dir("./images"))
	// session中间件
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store)) // 设置cookie名称

	publicRouter := router.Group("")
	{
		publicRouter.GET("/health", func(c *gin.Context) {
			c.JSON(200, "ok")
		})

		user := publicRouter.Group("user")
		{
			user.POST("/login", users.Login)
			user.POST("/register", users.Register)
		}

		articlePlu := router.Group("article")
		{
			articlePlu.POST("getArticle", articles.GetArticle)
		}
	}

	// session
	priviteRouter := router.Group("")
	priviteRouter.Use(middleware.SessVer())
	{
		// priviteRouter.POST("/test", api.Test)

		// 账号操作
		user := priviteRouter.Group("user")
		{
			user.POST("/logOut", users.LogOut)
			user.POST("/reset", users.Reset)
			user.POST("/checkEmail", users.CheckEmail)
			user.POST("/ResetPassword", users.ResetPassword)
			user.POST("/ChangPassword", users.ChangPassword)
		}

		// 文件上传
		uploade := priviteRouter.Group("upload")
		{
			uploade.POST("/uploadImage", uploads.Uploadfile_image)
		}

		// 文章
		article := priviteRouter.Group("article")
		{
			article.POST("/addArticle", articles.NewArticle)
			article.POST("/like", articles.Links)
		}
	}

	router.Run(r.port)
}

func CreateRouter() *Router {
	route := &Router{config.Configs.Port}
	// route.Run()
	return route
}
