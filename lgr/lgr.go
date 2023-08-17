package lgr

import "log"

var logger = log.Default()

func Print(msg string) {
	logger.Printf("[-- MY-log --]: %v\n", msg)
}
func Err(msg string) {
	logger.Printf("[## MY-BUG ##]: %v\n", msg)
}
