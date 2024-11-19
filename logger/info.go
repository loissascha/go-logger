package logger

import (
	"fmt"
)

func Info(err error, text string, vars ...any) {
	logErr(err)
	res := readTextVars(text)
	if len(res) != len(vars) {
		fmt.Println(FATAL_COLOR + "ATTENTION: Vars in string differ from provided vars" + RESET)
	}

	text = createTextOutput(text, res, vars)

	fmt.Println(INFO_COLOR+"INFO:", text+RESET)
}

func Debug(err error, text string, vars ...any) {
	logErr(err)
	res := readTextVars(text)
	if len(res) != len(vars) {
		fmt.Println(FATAL_COLOR + "ATTENTION: Vars in string differ from provided vars" + RESET)
	}

	text = createTextOutput(text, res, vars)

	fmt.Println(DEBUG_COLOR+"DEBUG:", text+RESET)
}

func logErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
