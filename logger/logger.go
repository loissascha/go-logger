package logger

type Logger struct {
	name          string
	color         string
	logToFile     bool
	showInConsole bool
}

func NewLogger(name string, color string, logToFile bool) *Logger {
	r := Logger{
		name:          name,
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
