package config

import "fmt"

func Load(path string) error {
	return nil
}

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
