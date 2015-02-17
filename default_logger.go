package log

import (
	"github.com/agtorre/gocolorize"
	"io"
	"log"
	"os"
)

func SetDefaultLogger() {
	Logger = &DefaultLogger{
		loggers: make(map[string]*log.Logger),
	}
}

type DefaultLogger struct {
	loggers map[string]*log.Logger
}

func (l *DefaultLogger) GetLogger(name string) *log.Logger {
	logger, ok := l.loggers[name]
	if !ok {
		switch name {
		case "fatal":
			logger = newColorLogger(os.Stderr, "FATAL", gocolorize.Blue)
		case "error":
			logger = newColorLogger(os.Stderr, "ERROR", gocolorize.Red)
		case "warn":
			logger = newColorLogger(os.Stderr, "WARN", gocolorize.Yellow)
		case "info":
			logger = newColorLogger(os.Stdout, "INFO", gocolorize.Cyan)
		case "debug":
			logger = newColorLogger(os.Stdout, "DEBUG", gocolorize.Green)
		}
		l.loggers[name] = logger
	}
	return logger
}

func newColorLogger(out io.Writer, severity string, color gocolorize.Color) *log.Logger {
	c := gocolorize.Colorize{Fg: color}
	return log.New(out, c.Paint("severity:"+severity)+"\t", log.Ldate|log.Ltime|log.Lshortfile)
}
