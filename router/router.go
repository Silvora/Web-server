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
		// api.GET("/status", func(ctx *gin.Context) {
		// 	str, _ := middleware.SetToken("757909414@qq.com")
		// 	ctx.JSON(http.StatusOK, gin.H{
		// 		"message": "Hello www.topgoer.com!",
		// 		"token":   str,
		// 	})

		// })
		api.GET("/login", service.Status)
		api.POST("/login", service.Login)
		api.POST("/sendEmail", service.UserToEmail)

		api.GET("/tag", service.GetTag)
		api.POST("/tag", service.AddTag)
		api.DELETE("/tag", service.DelTag)

		api.GET("/class", service.GetClass)
		api.POST("/class", service.AddClass)
		api.DELETE("/class", service.DelClass)

		api.GET("/markdown", service.GetMarkdown)
		api.POST("/markdown", service.AddMarkdown)
		api.PUT("/markdown", service.PutMarkdown)
		api.DELETE("/markdown", service.DelMarkdown)
	}

	return router
}
