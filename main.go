package main

import (
	"github.com/awesomexu/go-realworld/config"
	"github.com/awesomexu/go-realworld/middleware"
	"github.com/awesomexu/go-realworld/router"
	"github.com/awesomexu/go-realworld/utils"
	"github.com/awesomexu/go-realworld/validator"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)


func main() {
	config.InitConfig()
	utils.InitDB()

	r := gin.Default()

	middleware.LoadMiddleware(r)
	validator.RegisteMyValidator(r)
	router.LoadRouter(r)

	addr := viper.GetString("serverAddr")
	r.Run(addr)
}