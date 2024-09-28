package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql" // MySQL 驱动
)

var db *sql.DB

// 数据库连接初始化
func initDB() {
	var err error
	connStr := "root:123456@tcp(127.0.0.1:3306)/OJ"
	db, err = sql.Open("mysql", connStr)
	if err != nil {
		log.Fatal(err)
	}

	// 验证连接
	err = db.Ping()
	if err != nil {
		log.Fatal("Cannot connect to the database:", err)
	}
	fmt.Println("Connected to the database!")
}

// 查询数据库并返回结果的HTTP处理程序
func queryHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, name FROM mytable")
	if err != nil {
		http.Error(w, "Database query failed", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var result string
	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			http.Error(w, "Error scanning database results", http.StatusInternalServerError)
			return
		}
		result += fmt.Sprintf("ID: %d, Name: %s\n", id, name)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, "Error reading rows", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(result))
}

func main() {
	initDB()
	defer db.Close()

	// 设置HTTP路由
	http.HandleFunc("/query", queryHandler)

	// 启动HTTP服务器
	fmt.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
