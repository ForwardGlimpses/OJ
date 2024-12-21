package logs

import (
	"os"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func Init() {
	// 设置日志格式为JSON格式
	log.SetFormatter(&logrus.JSONFormatter{})

	// 打开日志文件，追加写入模式
    file, err := os.OpenFile("server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        logrus.Fatalf("Failed to open log file: %v", err)
    }
	log.SetOutput(file)

	// 设置日志输出到标准输出
	//log.SetOutput(os.Stdout)

	// 设置日志级别为Info级别
	log.SetLevel(logrus.InfoLevel)
}

func Info(args ...interface{}) {
	log.Info(args...)
}

func Infof(format string, args ...interface{}) {
	log.Infof(format, args...)
}

func Error(args ...interface{}) {
	log.Error(args...)
}

func Debug(args ...interface{}) {
	log.Debug(args...)
}

func Warn(args ...interface{}) {
	log.Warn(args...)
}
