package main

import (
	"todo_list/funcs"
	"todo_list/models"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化数据库
	models.InitDB()

	// 创建 Gin 实例
	r := gin.Default()

	// 注册 Todo 路由
	funcs.RegisterTodoRoutes(r)

	// 启动服务器
	r.Run(":8080")
}
