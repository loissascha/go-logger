package main

import (
	"fmt"
	"time"

	"github.com/loissascha/go-logger/logger"
)

func main() {
	// logger.Config.ShowDebug(false)

	logger.Config.AddFileLogging("/home/sascha/logs")

	someerr := fmt.Errorf("This is some weirdo errir")

	fmt.Println("Started...")
	timestamp := time.Now().UnixMilli()
	logger.Info(nil, "Some logging info text at time {timestamp} by user {username}", timestamp, "lois")
	logger.Debug(nil, "This is a debug message sent from {script}", "main")
	logger.Warning(someerr, "This is a warning!")
	logger.Error(nil, "This is an error!")
	logger.Fatal(nil, "This is a fatal!")

	customLogger := logger.NewLogger("CUSTOM", "", true)
	customLogger.Log(nil, "Some custom log text!")
	customLogger.Log(nil, "Some {var} with var", "text")
}
