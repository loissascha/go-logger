package logger

type config struct {
	showInfo    bool
	showDebug   bool
	showWarning bool
	showError   bool
	showFatal   bool
}

var Config config = config{
	showInfo:    true,
	showDebug:   true,
	showWarning: true,
	showError:   true,
	showFatal:   true,
}

func (c *config) ShowInfo(show bool) {
	c.showInfo = show
}

func (c *config) ShowDebug(show bool) {
	c.showDebug = show
}

func (c *config) ShowWarning(show bool) {
	c.showWarning = show
}

func (c *config) ShowError(show bool) {
	c.showError = show
}

func (c *config) ShowFatal(show bool) {
	c.showFatal = show
}
