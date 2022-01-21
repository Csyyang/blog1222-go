package router

import (
	"blog1222-go/config"

	"blog1222-go/api"

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

	// session中间件
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	publicRouter := router.Group("")
	{
		publicRouter.GET("/health", func(c *gin.Context) {
			c.JSON(200, "ok")
		})

		user := publicRouter.Group("user")
		{
			user.POST("/login", api.Login)
			user.POST("/register", api.Register)
			user.POST("/checkEmail", api.CheckEmail)
		}
	}

	priviteRouter := router.Group("")
	// 添加session验证
	priviteRouter.Use(middleware.SessVer())
	{
		priviteRouter.POST("/test", api.Test)

		user := priviteRouter.Group("user")
		user.POST("/logOut", api.LogOut)
		user.POST("/reset", api.Reset)

	}

	router.Run(r.port)
}

func CreateRouter() *Router {
	route := &Router{config.Configs.Port}
	// route.Run()
	return route
}
