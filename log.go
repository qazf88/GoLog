package golog

import (
	"fmt"
)

// ANSIC      = "Mon Jan _2 15:04:05 2006"
// UnixDate   = "Mon Jan _2 15:04:05 MST 2006"
// Stamp      = "Jan 02 15:04:05",
// RFC822Z    = "02 Jan 06 15:04 -0700",
// RFC3339    = "2006-01-02T15:04:05Z07:00",
// StampMilli = "Jan 02 15:04:05.000",
// StampMicro = "Jan 02 15:04:05.000000",
// StampNano  = "Jan _2 15:04:05.000000000".
func LogTimeFormat(format string) {
	timeFormat = format
}

// 0 = [TIME]   [APPNAME]  [LEVEL]   [/xxx/xx/main.go LINE]   "MESSAGE"
// 1 = [LEVEL]  [APPNAME]  [/xxx/xx/main.go LINE]   "MESSAGE"
// 2 = [LEVEL]  [APPNAME]  "MESSAGE"
// 3 = [LEVEL]  [APPNAME]  [TIME]    [/xxx/xx/main.go LINE]   "MESSAGE"
// 4 = [LEVEL]  [APPNAME]  [/xxx/xx/main.go LINE]   "MESSAGE"
func LogFormat(_format int) {
	_errorFormat = _format
	_warningFormat = _format
	_infoFormat = _format
	_debugFormat = _format
	_traceFormat = _format
	_fatalFormat = _format
}

// 1 = FATAL
// 2 = ERROR
// 3 = WARNING
// 4 = INFO
// 5 = DEBAG
// 6 = TRACE
func LogLevel(_level int) {
	logLevel = _level
}

//LogFatalFormat format FATAL message
func LogFatalFormat(_format int) {
	_fatalFormat = _format
}

//LogErrorFormat format ERROR message
func LogErrorFormat(_format int) {
	_errorFormat = _format
}

//LogWarningFormat format WARNING message
func LogWarningFormat(_format int) {
	_warningFormat = _format
}

//LogInfoFormat format INFO message
func LogInfoFormat(_format int) {
	_infoFormat = _format
}

//LogDebugFormat format DEBUG message
func LogDebugFormat(_format int) {
	_debugFormat = _format
}

//LogTraseFormat format TRACE message
func LogTraseFormat(_format int) {
	_traceFormat = _format
}

func Fatal(_message string) {
	if logLevel > 0 {
		fmt.Println(logBuild(_fatalFormat, _message, _fatal, red, true))
	}
}

func Error(_message string) {
	if logLevel > 1 {
		fmt.Println(logBuild(_errorFormat, _message, _error, red, false))
	}

}

func Warning(_message string) {
	if logLevel > 2 {
		fmt.Println(logBuild(_warningFormat, _message, _warning, yellow, false))
	}
}

func Info(_message string) {
	if logLevel > 3 {
		fmt.Println(logBuild(_infoFormat, _message, _info, green, false))
	}
}

func Debug(_message string) {
	if logLevel > 4 {
		fmt.Println(logBuild(_debugFormat, _message, _debug, cyan, false))
	}
}

func Trace(_message string) {
	if logLevel > 5 {
		fmt.Println(logBuild(_traceFormat, _message, _trace, gray, false))
	}
}

func LogVersion() string {
	return version
}
