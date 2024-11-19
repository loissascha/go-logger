package logger

import "fmt"

type VarInfo struct {
	Name       string
	StartIndex int
	StopIndex  int
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
