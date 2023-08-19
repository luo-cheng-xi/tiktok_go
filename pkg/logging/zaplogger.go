package logging

import (
	"go.uber.org/zap"
)

var Logger = loadDebug()

func loadDebug() *zap.Logger {
	l, err := zap.NewProduction()
	if err != nil {
		panic("日志对象生成异常")
	}
	return l
}
