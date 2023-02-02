package main

import (
	"indexServer/config"
	"indexServer/db"
	"indexServer/logger"
	"indexServer/router"
)

func main() {
	//开启log日志
	logger.InitLogger()

	//开启Mysql
	db.InitMysql()

	//开启Redis
	db.InitRedis()

	//开启邮箱验证
	//middleware.InitEmail()

	//middleware.SendEmail("757909414@qq.com")

	//开启路由
	route := router.InitRouters()

	//开启端口
	route.Run(config.Port)

}
