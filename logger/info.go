package logger

import (
	"fmt"
)

func Info(err error, text string, vars ...any) {
	logErr(err)
	res := readTextVars(text)
	resErr(res, vars)

	if Config.showInfo {
		text = createTextOutput(text, res, vars)
		fmt.Println(color_info+"INFO:", text+color_reset)
	}
}

func Debug(err error, text string, vars ...any) {
	logErr(err)
	res := readTextVars(text)
	resErr(res, vars)

	if Config.showDebug {
		text = createTextOutput(text, res, vars)
		fmt.Println(color_debug+"DEBUG:", text+color_reset)
	}
}

func Warning(err error, text string, vars ...any) {
	logErr(err)
	res := readTextVars(text)
	resErr(res, vars)

	if Config.showWarning {
		text = createTextOutput(text, res, vars)
		fmt.Println(color_warning+"WARNING:", text+color_reset)
	}
}

func Error(err error, text string, vars ...any) {
	logErr(err)
	res := readTextVars(text)
	resErr(res, vars)

	if Config.showError {
		text = createTextOutput(text, res, vars)
		fmt.Println(color_error+"ERROR:", text+color_reset)
	}
}

func Fatal(err error, text string, vars ...any) {
	logErr(err)
	res := readTextVars(text)
	resErr(res, vars)

	if Config.showFatal {
		text = createTextOutput(text, res, vars)
		fmt.Println(color_fatal+"FATAL:", text+color_reset)
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
