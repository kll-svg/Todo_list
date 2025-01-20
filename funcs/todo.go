package funcs

import (
	"net/http"
	"todo_list/models"

	"github.com/gin-gonic/gin"
)

// 注册 Todo 路由
func RegisterTodoRoutes(r *gin.Engine) {
	todoGroup := r.Group("/todos")
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

	if err := models.DB.Create(&todo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create todo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Todo created successfully", "data": todo})
}

// 获取所有任务
func GetTodos(c *gin.Context) {
	var todos []models.Todo
	if err := models.DB.Find(&todos).Error; err != nil {
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
	if err := models.DB.Delete(&models.Todo{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete todo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully"})
}
