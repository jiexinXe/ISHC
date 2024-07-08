package main

import (
	"ISHC/config"
	_ "ISHC/docs" // 必须导入 Swag 生成的 docs 包
	"ISHC/middleware"
	"ISHC/router"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title ISHC API
// @version 1.0
// @description This is a sample server for ISHC project.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /

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

	// Swagger 路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 启动服务器
	r.Run(":8080")
}
