package router

import (
	"blog1222-go/config"

	"github.com/gin-gonic/gin"
)

type Router struct {
	port string
}

func (r *Router) Run() {
	router := gin.Default()
	router.Run(r.port)
}

func CreateRouter() *Router {
	route := &Router{config.Configs.Port}
	// route.Run()
	return route
}
