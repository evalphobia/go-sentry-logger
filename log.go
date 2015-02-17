/*
	this package enable logging data with sending to sentry.
	see: https://github.com/evalphobia/go-sentry-logger
*/

package log

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"runtime"
)

type LogData struct {
	Request *http.Request
	Label   string

	data  Any
	depth int
	err   error
}

func NewLogData(data Any, depth int) *LogData {
	return &LogData{
		data:  data,
		depth: depth,
	}
}

// logging fatal error
func Fatal(l *LogData) {
	logger := Logger.GetLogger("fatal")
	file, line := getTrace(l.depth)
	toSentry := sentryLevel >= sentryLevelFatal
	logging(logger, l, file, line, toSentry)
}

// logging high level error
func Error(l *LogData) {
	logger := Logger.GetLogger("error")
	file, line := getTrace(l.depth)
	toSentry := sentryLevel >= sentryLevelError
	logging(logger, l, file, line, toSentry)
}

// logging middle level error
func Warn(l *LogData) {
	logger := Logger.GetLogger("warn")
	file, line := getTrace(l.depth)
	toSentry := sentryLevel >= sentryLevelWarn
	logging(logger, l, file, line, toSentry)
}

// logging infomation
func Info(l *LogData) {
	logger := Logger.GetLogger("info")
	file, line := getTrace(l.depth)
	toSentry := sentryLevel >= sentryLevelInfo
	logging(logger, l, file, line, toSentry)
}

// logging debug infomation
func Debug(l *LogData) {
	logger := Logger.GetLogger("debug")
	file, line := getTrace(l.depth)
	toSentry := sentryLevel >= sentryLevelDebug
	logging(logger, l, file, line, toSentry)
}

// logging error without sentry
func ErrorWithoutTrack(l *LogData) {
	logger := Logger.GetLogger("error")
	file, line := getTrace(l.depth)
	logging(logger, l, file, line, false)
}

// print infomation with the logging format
func Print(v Any, depth int) {
	l := NewLogData(v, depth)
	file, line := getTrace(l.depth)
	m := composeLogData(file, line, l)
	printData(m)
}

// print header
func PrintHeader(v ...Any) {
	var val string
	if len(v) > 0 {
		val = fmt.Sprint(v[0])
	} else {
		val = ""
	}
	fmt.Printf("\n============================ %s ============================\n", val)
}

// write log, send sentry
func logging(logger *log.Logger, l *LogData, file string, line int, toSentry bool) {
	mapData := composeLogData(file, line, l)
	writeLog(logger, mapData)
	if mapData["type"] == "error" {
		l.err = l.data.(error)
	} else {
		l.err = errors.New("general error")
	}
	if toSentry {
		SendSentryMapData(mapData, l.Request, l.err)
	}
}

// write log to logger
func writeLog(l *log.Logger, m map[string]string) {
	format := "label:%s\tfile:%s\tline:%s\ttype:%s\tdata:%s\taddr:%s"
	l.Printf(format, m["label"], m["file"], m["line"], m["type"], m["data"], m["addr"])
}

// print data to stdout
func printData(m map[string]string) {
	format := "label:%s\tfile:%s\tline:%s\ttype:%s\tdata:%s\taddr:%s\n"
	fmt.Printf(format, m["label"], m["file"], m["line"], m["type"], m["data"], m["addr"])
}

// conpose log data from filename, line, data
func composeLogData(file string, line int, l *LogData) map[string]string {
	m := make(map[string]string)
	m["label"] = l.Label
	m["file"] = file
	m["line"] = ParseToString(line)
	m["type"] = ParseToType(l.data)
	m["data"] = ParseToString(l.data)
	if HasPointer(m["type"]) {
		m["addr"] = fmt.Sprintf("%p", l.data)
	} else {
		m["addr"] = "(none)"
	}
	return m
}

// return stack trace infomation
func getTrace(num int) (string, int) {
	_, file, line, ok := runtime.Caller(num)
	if !ok {
		file = "(cannot trace file)"
		line = 0
	} else {
		short := file
		for i := len(file) - 1; i > 0; i-- {
			if file[i] == '/' {
				short = file[i+1:]
				break
			}
		}
		file = short
	}
	return file, line
}
