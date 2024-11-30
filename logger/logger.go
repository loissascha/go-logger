package logger

type Logger struct {
	prefix          string
	color         string
	logToFile     bool
	showInConsole bool
}

func NewLogger(prefix string, color string, logToFile bool) *Logger {
	r := Logger{
		prefix:          prefix,
		color:         color,
		logToFile:     logToFile,
		showInConsole: true,
	}
	return &r
}

func (l *Logger) ShowInConsole(show bool) {
	l.showInConsole = show
}

func (l *Logger) Log(t string) {

}
