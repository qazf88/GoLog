package golog

import (
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"time"
)

// Build log message
func logBuild(format int, _message string, _level string, messageColor string, fatal bool, _line bool) string {
	color := white
	if fatal {
		color = red
	}
	switch format {
	case 0:
		return logTime(timeFormat) + _level + color + appName() + reset + color + line(_line) + reset + messageColor + _message + reset
	case 1:
		return _level + color + appName() + line(_line) + reset + messageColor + _message + reset
	case 2:
		return _level + color + appName() + reset + messageColor + _message + reset
	case 3:
		return _level + color + appName() + reset + logTime(timeFormat) + color + line(_line) + reset + messageColor + _message + reset
	case 4:
		return _level + color + appName() + line(_line) + reset + messageColor + _message + reset
	default:
		return logTime(timeFormat) + color + appName() + reset + _level + color + line(_line) + reset + messageColor + _message + reset
	}
}

func logTime(format string) string {
	res := time.Now().Format(format)
	return green + "[" + res + "] " + reset
}

func line(_line bool) string {
	_, file, line, _ := runtime.Caller(3)
	if _line {
		return "[" + file + ":" + strconv.Itoa(line) + "]  "
	}
	return "[" + file + "]  "
}

func appName() string {
	return "[" + filepath.Base(os.Args[0]) + "] "
}
