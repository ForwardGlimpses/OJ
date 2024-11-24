package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func Load(path string) error {
	//add 读configs.json 并解析到C里面
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("could not open config file: %v", err)
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("could not read config file: %v", err)
	}

	var cfg Config
	if err := json.Unmarshal(bytes, &cfg); err != nil {
		return fmt.Errorf("could not parse config file: %v", err)
	}
	C = cfg
	return nil
}

//启动 go run "C:\Users\乔书祥\Desktop\OJ\server\cmd\main.go" start -c C:\Users\乔书祥\Desktop\OJ\server\configs\configs.json

var C Config

type Config struct {
	Mysql Mysql
}

type Mysql struct {
	User     string
	Password string
	Host     string
	Port     int
	DBName   string
}

func (a Mysql) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		a.User, a.Password, a.Host, a.Port, a.DBName)
}
