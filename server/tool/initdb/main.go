package main

import (
	"github.com/ForwardGlimpses/OJ/server/pkg/logs"
	"github.com/ForwardGlimpses/OJ/server/pkg/schema"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// 连接到 MySQL 数据库
	dsn := "root:111111@tcp(127.0.0.1:3306)/ojmysql?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logs.Error("连接数据库失败:", err)
		return
	}

	// 自动迁移：创建表结构
	err = db.AutoMigrate(&schema.UsersDBItem{}, &schema.ContestDBItem{}, &schema.ContestSolutionDBItem{}, &schema.ContestProblemDBItem{}, &schema.ProblemDBItem{}, &schema.SolutionDBItem{}, &schema.SourceCodeDBItem{})
	if err != nil {
		logs.Error("创建表失败:", err)
		return
	}

	logs.Info("表创建成功！")
}
