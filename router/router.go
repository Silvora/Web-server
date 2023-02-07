package router

import (
	"indexServer/middleware"
	"indexServer/service"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func serverLog() {
	gin.DisableConsoleColor()
	file, _ := os.Create("log/gin.log")
	//gin.DefaultWriter = io.MultiWriter(file)
	gin.DefaultWriter = io.MultiWriter(file, os.Stdout)
}

func InitRouters() *gin.Engine {

	serverLog()

	router := gin.Default()
	router.Use(middleware.Cors())
	router.Use(middleware.GetTokenAuth())

	api := router.Group("/api")
	{
		//登录注册
		api.GET("/login", service.Status)
		api.POST("/login", service.Login)
		api.POST("/sendEmail", service.UserToEmail)
	}

	//blog
	blog := router.Group("/blog")
	{

		//blog
		blog.GET("/tag", service.GetTag)
		blog.POST("/tag", service.AddTag)
		blog.DELETE("/tag", service.DelTag)

		blog.GET("/class", service.GetClass)
		blog.POST("/class", service.AddClass)
		blog.DELETE("/class", service.DelClass)

		blog.GET("/markdown", service.GetMarkdown)
		blog.POST("/markdown", service.AddMarkdown)
		blog.PUT("/markdown", service.PutMarkdown)
		blog.DELETE("/markdown", service.DelMarkdown)
	}

	return router
}
