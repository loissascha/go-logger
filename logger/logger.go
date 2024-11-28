package logger

type Logger struct {
	name             string
	color            string
	logToFile        bool
	showInConsole    bool
	fileLoggingAsync bool
}

func NewLogger(name string, color string, logToFile bool) {

}

func (l *Logger) Log(t string) {

}
