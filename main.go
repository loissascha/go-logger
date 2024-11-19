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
}
