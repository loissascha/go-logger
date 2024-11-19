package logger

import (
	"fmt"
)

func Info(err error, text string, vars ...any) {
	if err != nil {
		fmt.Println(err)
	}
	res := readTextVars(text)
	if len(res) != len(vars) {
		fmt.Println("ATTENTION: Vars in string differ from provided vars")
	}

	text = createTextOutput(text, res, vars)

	fmt.Println(INFO_COLOR+"INFO:", text+RESET)
}
