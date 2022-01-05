package router

import (
	"blog1222-go/config"

	"blog1222-go/api"

	"github.com/gin-gonic/gin"

	"blog1222-go/middleware"
)

type Router struct {
	port string
}

func (r *Router) Init() {
	router := gin.Default()

	publicRouter := router.Group("")
	{
		publicRouter.GET("/health", func(c *gin.Context) {
			c.JSON(200, "ok")
		})

		user := publicRouter.Group("user")
		{
			user.POST("/login", api.Login)
			user.POST("/register", api.Register)
		}
	}

	priviteRouter := router.Group("")
	priviteRouter.Use(middleware.JwtVer())
	{
		priviteRouter.GET("/test")
	}

	router.Run(r.port)
}

func CreateRouter() *Router {
	route := &Router{config.Configs.Port}
	// route.Run()
	return route
}
