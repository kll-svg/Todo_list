package models

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// 初始化数据库
func InitDB() {
	// 数据库连接配置
	dsn := "root:ZhengZhiXing719@tcp(127.0.0.1:3306)/tododb?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// 自动迁移表结构
	err = DB.AutoMigrate(&Todo{}, &User{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
}

// Todo 数据模型
type Todo struct {
	ID      int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Content string `json:"content"`
	Done    bool   `json:"done"`
	UserID  int    `json:"user_id"` // 关联到用户的 ID
}

// User 数据模型
type User struct {
	ID       int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Username string `json:"username" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
}
