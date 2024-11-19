package main

import (
	"fmt"
	"time"

	"github.com/loissascha/go-logger/logger"
)

func main() {
	fmt.Println("Started...")
	timestamp := time.Now().UnixMilli()
	logger.Info(nil, "Some logging text with {timestamp} and {username}", timestamp, "lois")
	fmt.Println("after info")
	logger.Debug(nil, "This is a debug message sent from {script}", "main")
	logger.Warning(nil, "This is a warning!")
	logger.Error(nil, "This is an error!")
	logger.Fatal(nil, "This is a fatal!")
}
