package lgr

import "log"

var logger = log.Default()

func Print(msg string) {
	logger.Printf("[-- MY-log --]: %v\n", msg)
}
