package global

import (
	"net/http"

	"github.com/ForwardGlimpses/OJ/server/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB         *gorm.DB
	HttpClient = &http.Client{}
)

func Init() error {
	var err error
	DB, err = gorm.Open(mysql.Open(config.C.Mysql.DSN()), &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}
