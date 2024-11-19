package logger

import (
	"fmt"
)

type VarInfo struct {
	Name       string
	StartIndex int
	StopIndex  int
}

const (
	INFO_COLOR = "\033[34m"
	DEBUG_COLOR = "\033[32m"
	WARNING_COLOR = "\033[33m"
	ERROR_COLOR = "\033[31m"
	FATAL_COLOR = "\033[35m"
	RESET = "\033[0m"
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

	fmt.Println(INFO_COLOR + "INFO:", text + RESET)
}

func createTextOutput(text string, varInfo []VarInfo, vars []any) string {
	for i := len(varInfo) - 1; i >= 0; i-- {
		v := varInfo[i]
		before := text[0:v.StartIndex]
		after := text[v.StopIndex+1:]
		text = before + fmt.Sprintf("%v", vars[i]) + after
	}
	return text
}

func readTextVars(text string) []VarInfo {
	readVar := false
	currentVar := ""
	currentVarStartIndex := 0
	vars := []VarInfo{}
	for i, v := range text {
		char := string(v)

		if char == "}" {
			readVar = false
			vars = append(vars, VarInfo{
				Name:       currentVar,
				StartIndex: currentVarStartIndex,
				StopIndex:  i,
			})
			currentVar = ""
		}

		if readVar {
			currentVar += char
		}

		if char == "{" {
			readVar = true
			currentVarStartIndex = i
		}
	}
	return vars
}
