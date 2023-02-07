package db

import (
	"indexServer/config"
	"indexServer/logger"
	"log"

	"github.com/garyburd/redigo/redis"
)

var Pool *redis.Pool

func InitRedis() {
	idle := config.RedisMaxIdle
	active := config.RedisMaxActive
	host := config.RedisIP
	port := config.RedisPort

	Pool = &redis.Pool{
		MaxIdle:     idle,   /*最大的空闲连接数*/
		MaxActive:   active, /*最大的激活连接数*/
		IdleTimeout: 300,
		Dial: func() (redis.Conn, error) {
			Conn, err := redis.Dial("tcp", host+":"+port)
			if err != nil {
				logger.SetLogger(1, "redis连接失败")
				log.Println("redis连接失败")
				return nil, err
			}
			// if _, err := Conn.Do("AUTH", 757909414); err != nil {
			// 	logger.SetLogger(1, "redis连接密码失败")
			// 	log.Println("redis连接密码失败")
			// 	Conn.Close()
			// 	return nil, err
			// }

			logger.SetLogger(0, "redis连接成功")
			log.Println("redis连接成功")
			return Conn, nil
		},
	}
}

func SetRedis(key string, value string, timeout string) bool {
	conn := Pool.Get()
	defer conn.Close()

	_, err := conn.Do("SET", key, value, "EX", timeout)
	if err != nil {
		logger.SetLogger(1, "Redis插入出错了")
		log.Println("插入出错了", err)
		return false
	}
	//log.Println("Redis插入成功")
	//middleware.SetLogger(0, "Redis插入成功")
	return true
}

func GetRedis(key string) string {
	conn := Pool.Get()
	defer conn.Close()

	value, err := redis.String(conn.Do("GET", key))
	if err != nil {
		logger.SetLogger(1, "Redis插入出错了")
		log.Println("插入出错了", err)
		return ""
	}
	return value
}
