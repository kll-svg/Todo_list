package funcs

import (
	"net/http"
	"todo_list/models"

	"github.com/gin-gonic/gin"
)

// 注册 Todo 路由
func RegisterTodoRoutes(routerGroup *gin.RouterGroup) {
	todoGroup := routerGroup.Group("/todos")
	{
		todoGroup.POST("/", CreateTodo)      // 新增任务
		todoGroup.GET("/", GetTodos)         // 获取所有任务
		todoGroup.PUT("/:id", UpdateTodo)    // 更新任务
		todoGroup.DELETE("/:id", DeleteTodo) // 删除任务
	}
}

// 新增任务
func CreateTodo(c *gin.Context) {
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	// 从上下文获取用户 ID
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}
	todo.UserID = userID.(int)

	if err := models.DB.Create(&todo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create todo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Todo created successfully", "data": todo})
}

// 获取所有任务
func GetTodos(c *gin.Context) {
	var todos []models.Todo

	// 从上下文获取用户 ID
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// 查询当前用户的任务
	if err := models.DB.Where("user_id = ?", userID).Find(&todos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch todos"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": todos})
}

// 更新任务
func UpdateTodo(c *gin.Context) {
	id := c.Param("id")
	var todo models.Todo

	// 查找任务
	if err := models.DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	// 检查任务是否属于当前用户
	userID, exists := c.Get("userID")
	if !exists || todo.UserID != userID.(int) {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized to update this todo"})
		return
	}

	// 绑定更新数据
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	// 更新任务
	if err := models.DB.Save(&todo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update todo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Todo updated successfully", "data": todo})
}

// 删除任务
func DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	var todo models.Todo

	// 查找任务
	if err := models.DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	// 检查任务是否属于当前用户
	userID, exists := c.Get("userID")
	if !exists || todo.UserID != userID.(int) {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized to delete this todo"})
		return
	}

	// 删除任务
	if err := models.DB.Delete(&todo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete todo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully"})
}
