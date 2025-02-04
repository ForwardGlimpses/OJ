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

//启动 go run ".\cmd\main.go" start -c .\configs\configs.json

var C Config

type Config struct {
	Mysql Mysql
	Judge JudgeConfig
	Root  Root
}

type Mysql struct {
	User     string
	Password string
	Host     string
	Port     int
	DBName   string
	Debug    bool
}

type JudgeConfig struct {
	Host string
	Port int
}

type Root struct {
	Email    string
	Password string
}

func (a Mysql) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		a.User, a.Password, a.Host, a.Port, a.DBName)
}

func (j JudgeConfig) BaseURL() string {
	return fmt.Sprintf("http://%s:%d/run", j.Host, j.Port)
}
