package logger

import "fmt"

type Logger struct {
	prefix        string
	color         string
	logToFile     bool
	showInConsole bool
}

func NewLogger(prefix string, color string, logToFile bool) *Logger {
	r := Logger{
		prefix:        prefix,
		color:         color,
		logToFile:     logToFile,
		showInConsole: true,
	}
	return &r
}

func (l *Logger) ShowInConsole(show bool) {
	l.showInConsole = show
}

func (l *Logger) Log(err error, text string, vars ...any) {
	// TODO: logErr

	res := readTextVars(text)
	resErr(res, vars)
	text = createTextOutput(text, res, vars)
	ts := getTimeString()

	if l.showInConsole {
		fmt.Println(l.color+ts+l.prefix+":", text+color_reset)
	}

	// TODO: file logging!
}
