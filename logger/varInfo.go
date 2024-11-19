package logger

import "fmt"

type varInfo struct {
	Name       string
	StartIndex int
	StopIndex  int
}

func createTextOutput(text string, varInfo []varInfo, vars []any) string {
	for i := len(varInfo) - 1; i >= 0; i-- {
		v := varInfo[i]
		before := text[0:v.StartIndex]
		after := text[v.StopIndex+1:]
		text = before + fmt.Sprintf("%v", vars[i]) + after
	}
	return text
}

func readTextVars(text string) []varInfo {
	readVar := false
	currentVar := ""
	currentVarStartIndex := 0
	vars := []varInfo{}
	for i, v := range text {
		char := string(v)

		if char == "}" {
			readVar = false
			vars = append(vars, varInfo{
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
