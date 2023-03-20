package router

import (
	"github.com/awesomexu/go-realworld/controller"
	"github.com/awesomexu/go-realworld/middleware"
	"github.com/gin-gonic/gin"
)

func LoadRouter(r *gin.Engine){
	api := r.Group("/api")
	LoadUserRouter(api)
}

func LoadUserRouter(r *gin.RouterGroup){
	r.POST("/users", controller.Register)
	r.POST("/users/login", controller.Login)
	r.GET("/user", middleware.Auth, controller.GetCurrentUser)
	r.PUT("/user", middleware.Auth, controller.UpdateUserInfo)
}

