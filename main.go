package main

import (
	"ISHC/config"
	"ISHC/middleware"
	"ISHC/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 使用CORS中间件
	r.Use(middleware.CORS())

	// 初始化数据库
	config.InitDB()

	// 初始化 Kafka
	config.InitKafka()

	// 初始化路由
	router.InitRoutes(r)

	// 启动服务器
	r.Run(":8080")
}
