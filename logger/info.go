package logger

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type LoggingType int

const (
	LOG_INFO LoggingType = iota
	LOG_DEBUG
	LOG_WARNING
	LOG_ERROR
	LOG_FATAL
)

func getTimeString() string {
	ts := ""
	if Config.printDatesInConsole {
		t := time.Now()
		tf := t.Format("06-01-02 15:04:05")
		ts = fmt.Sprintf("[%v] ", tf)
	}
	return ts
}

func logToLogFiles(lt LoggingType, text string) {
	for _, v := range Config.logPaths {
		file := strings.TrimSuffix(v, "/")
		file += "/all.log"

		f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println("Can't write file", file)
			continue
		}
		defer f.Close()

		_, err = f.WriteString(text)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			continue
		}
	}
}

func Info(err error, text string, vars ...any) {
	logErr(LOG_INFO, err)
	res := readTextVars(text)
	resErr(res, vars)
	text = createTextOutput(text, res, vars)
	ts := getTimeString()

	if Config.showInfo {
		fmt.Println(color_info+ts+"INFO:", text+color_reset)
	}

	logToLogFiles(LOG_INFO, ts+"INFO: "+text+"\n")
}

func Debug(err error, text string, vars ...any) {
	logErr(LOG_DEBUG, err)
	res := readTextVars(text)
	resErr(res, vars)
	text = createTextOutput(text, res, vars)
	ts := getTimeString()

	if Config.showDebug {
		fmt.Println(color_debug+ts+"DEBUG:", text+color_reset)
	}

	logToLogFiles(LOG_DEBUG, ts+"DEBUG: "+text+"\n")
}

func Warning(err error, text string, vars ...any) {
	logErr(LOG_WARNING, err)
	res := readTextVars(text)
	resErr(res, vars)
	text = createTextOutput(text, res, vars)
	ts := getTimeString()

	if Config.showWarning {
		fmt.Println(color_warning+ts+"WARNING:", text+color_reset)
	}

	logToLogFiles(LOG_WARNING, ts+"WARNING: "+text+"\n")
}

func Error(err error, text string, vars ...any) {
	logErr(LOG_ERROR, err)
	res := readTextVars(text)
	resErr(res, vars)
	text = createTextOutput(text, res, vars)
	ts := getTimeString()

	if Config.showError {
		fmt.Println(color_error+ts+"ERROR:", text+color_reset)
	}

	logToLogFiles(LOG_ERROR, ts+"ERROR: "+text+"\n")
}

func Fatal(err error, text string, vars ...any) {
	logErr(LOG_FATAL, err)
	res := readTextVars(text)
	resErr(res, vars)
	text = createTextOutput(text, res, vars)
	ts := getTimeString()

	if Config.showFatal {
		fmt.Println(color_fatal+ts+"FATAL:", text+color_reset)
	}

	logToLogFiles(LOG_FATAL, ts+"FATAL: "+text+"\n")
}

func logErr(lt LoggingType, err error) {
	if err == nil {
		return
	}
	ts := getTimeString()
	switch lt {
	case LOG_INFO:
		fmt.Print(color_info + ts + "INFO:")
		break
	case LOG_DEBUG:
		fmt.Print(color_debug + ts + "DEBUG:")
		break
	case LOG_WARNING:
		fmt.Print(color_warning + ts + "WARNING:")
		break
	case LOG_ERROR:
		fmt.Print(color_error + ts + "ERROR:")
		break
	case LOG_FATAL:
		fmt.Print(color_fatal + ts + "FATAL:")
		break
	}
	fmt.Print(err)
	fmt.Print(color_reset, "\n")
}

func resErr(res []varInfo, vars []any) {
	if len(res) != len(vars) {
		fmt.Println(color_fatal + "ATTENTION: Vars in string differ from provided vars" + color_reset)
	}
}
