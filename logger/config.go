package logger

import (
	"fmt"
	"os"
)

type config struct {
	showInfo            bool
	showDebug           bool
	showWarning         bool
	showError           bool
	showFatal           bool
	fileLogInfo         bool
	fileLogDebug        bool
	fileLogWarning      bool
	fileLogError        bool
	fileLogFatal        bool
	printDatesInConsole bool
	logPaths            []string
	fileLoggingAsync    bool
}

var Config config = config{
	showInfo:            true,
	showDebug:           true,
	showWarning:         true,
	showError:           true,
	showFatal:           true,
	fileLogInfo:         true,
	fileLogDebug:        false,
	fileLogWarning:      true,
	fileLogError:        true,
	fileLogFatal:        true,
	printDatesInConsole: true,
	logPaths:            []string{},
	fileLoggingAsync:    true,
}

func (c *config) FileLogInfo(log bool) {
	c.fileLogInfo = log
}

func (c *config) FileLogDebug(log bool) {
	c.fileLogDebug = log
}

func (c *config) FileLogWarning(log bool) {
	c.fileLogWarning = log
}

func (c *config) FileLogError(log bool) {
	c.fileLogError = log
}

func (c *config) FileLogFatal(log bool) {
	c.fileLogFatal = log
}

func (c *config) ShowInfo(show bool) {
	c.showInfo = show
}

func (c *config) ShowDebug(show bool) {
	c.showDebug = show
}

func (c *config) ShowWarning(show bool) {
	c.showWarning = show
}

func (c *config) ShowError(show bool) {
	c.showError = show
}

func (c *config) ShowFatal(show bool) {
	c.showFatal = show
}

func (c *config) PrintDates(show bool) {
	c.printDatesInConsole = show
}

func (c *config) AddFileLogging(path string) {
	isDir, err := isDirectory(path)
	if err != nil || !isDir {
		Error(err, "AddFileLogging failed. Path {path} is not a directory.", path)
		return
	}
	c.logPaths = append(c.logPaths, path)
}

func (c *config) AsyncFileLogging(async bool) {
	c.fileLoggingAsync = async
}

func isDirectory(path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, fmt.Errorf("path does not exist")
		}
		return false, err
	}
	return info.IsDir(), nil
}
