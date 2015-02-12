package log

import (
	"github.com/agtorre/gocolorize"
	"io"
	"log"
	"os"
)

func SetDefaultLogger(loggers map[string]*log.Logger) {
	loggers["fatal"] = newColorLogger(os.Stderr, "FATAL", gocolorize.Blue)
	loggers["error"] = newColorLogger(os.Stderr, "ERROR", gocolorize.Red)
	loggers["warn"] = newColorLogger(os.Stderr, "WARN", gocolorize.Yellow)
	loggers["info"] = newColorLogger(os.Stdout, "INFO", gocolorize.Cyan)
	loggers["debug"] = newColorLogger(os.Stdout, "DEBUG", gocolorize.Green)
}

func newColorLogger(out io.Writer, severity string, color gocolorize.Color) *log.Logger {
	c := gocolorize.Colorize{Fg: color}
	return log.New(out, c.Paint("severity:"+severity)+"\t", log.Ldate|log.Ltime|log.Lshortfile)
}
