package logger

import (
	"fmt"
	"time"
)

func Info(err error, text string, vars ...any) {
	logErr(err)
	res := readTextVars(text)
	resErr(res, vars)
	text = createTextOutput(text, res, vars)

	if Config.showInfo {
		ts := ""
		if Config.printDatesInConsole {
			t := time.Now()
			tf := t.Format("06-01-02 15:04:05")
			ts = fmt.Sprintf("[%v] ", tf)
		}
		fmt.Println(color_info+ts+"INFO:", text+color_reset)
	}
}

func Debug(err error, text string, vars ...any) {
	logErr(err)
	res := readTextVars(text)
	resErr(res, vars)
	text = createTextOutput(text, res, vars)

	if Config.showDebug {
		ts := ""
		if Config.printDatesInConsole {
			t := time.Now()
			tf := t.Format("06-01-02 15:04:05")
			ts = fmt.Sprintf("[%v] ", tf)
		}
		fmt.Println(color_debug+ts+"DEBUG:", text+color_reset)
	}
}

func Warning(err error, text string, vars ...any) {
	logErr(err)
	res := readTextVars(text)
	resErr(res, vars)
	text = createTextOutput(text, res, vars)

	if Config.showWarning {
		ts := ""
		if Config.printDatesInConsole {
			t := time.Now()
			tf := t.Format("06-01-02 15:04:05")
			ts = fmt.Sprintf("[%v] ", tf)
		}
		fmt.Println(color_warning+ts+"WARNING:", text+color_reset)
	}
}

func Error(err error, text string, vars ...any) {
	logErr(err)
	res := readTextVars(text)
	resErr(res, vars)
	text = createTextOutput(text, res, vars)

	if Config.showError {
		ts := ""
		if Config.printDatesInConsole {
			t := time.Now()
			tf := t.Format("06-01-02 15:04:05")
			ts = fmt.Sprintf("[%v] ", tf)
		}
		fmt.Println(color_error+ts+"ERROR:", text+color_reset)
	}
}

func Fatal(err error, text string, vars ...any) {
	logErr(err)
	res := readTextVars(text)
	resErr(res, vars)
	text = createTextOutput(text, res, vars)

	if Config.showFatal {
		ts := ""
		if Config.printDatesInConsole {
			t := time.Now()
			tf := t.Format("06-01-02 15:04:05")
			ts = fmt.Sprintf("[%v] ", tf)
		}
		fmt.Println(color_fatal+ts+"FATAL:", text+color_reset)
	}
}

func logErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func resErr(res []varInfo, vars []any) {
	if len(res) != len(vars) {
		fmt.Println(color_fatal + "ATTENTION: Vars in string differ from provided vars" + color_reset)
	}
}
