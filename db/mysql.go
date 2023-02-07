package db

import (
	"database/sql"
	"indexServer/config"
	"indexServer/logger"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var MysqlDB *sql.DB

func InitMysql() {
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	name := config.MysqlName
	pass := config.MysqlPass
	ip := config.MysqlIP
	prot := config.MysqlPort
	db := config.MysqlDB
	dns := name + ":" + pass + "@tcp(" + ip + ":" + prot + ")/" + db + "?charset=utf8"
	//fmt.Println(dns)

	//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	MysqlDB, _ = sql.Open("mysql", dns)

	//设置数据库最大连接数
	MysqlDB.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	MysqlDB.SetMaxIdleConns(10)

	//验证连接
	if err := MysqlDB.Ping(); err != nil {
		log.Println("mysql连接失败")
		logger.SetLogger(1, "mysql连接失败")
		return
	}

	log.Println("mysql连接成功")
	logger.SetLogger(0, "mysql连接成功")
}
