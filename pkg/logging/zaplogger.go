package logging

import (
	"go.uber.org/zap"
)

var logger = loadDebug()

func NewLogger() *zap.Logger {
	return logger
}

func loadDebug() *zap.Logger {
	l, err := zap.NewProduction()
	if err != nil {
		panic("日志对象生成异常")
	}
	return l
}
