package main

import (
	"todo_list/auth"
	"todo_list/funcs"
	"todo_list/models"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化数据库
	models.InitDB()

	// 创建 Gin 实例
	r := gin.Default()

	// 用户登录注册路由
	r.POST("/register", auth.Register)
	r.POST("/login", auth.Login)

	// 任务路由（需要鉴权）
	authorized := r.Group("/")
	authorized.Use(auth.Authenticate())
	funcs.RegisterTodoRoutes(authorized) // 传入 *gin.RouterGroup

	// 启动服务器
	r.Run(":8080")
}
