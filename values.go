package golog

var (
	logLevel       int = 0
	timeFormat         = UnixDate
	_fatalFormat       = 0
	_errorFormat       = 0
	_warningFormat     = 0
	_infoFormat        = 2
	_debugFormat       = 0
	_traceFormat       = 0
)

const (
	version = "v0.0.4"

	_fatal   = red + "[FATAL]   " + reset
	_error   = red + "[ERROR]   " + reset
	_warning = yellow + "[WARN]    " + reset
	_info    = green + "[INFO]    " + reset
	_debug   = cyan + "[DEBUG]   " + reset
	_trace   = gray + "[TRACE]   " + reset

	reset  = "\033[0m"
	red    = "\033[31m"
	green  = "\033[32m"
	yellow = "\033[33m"
	blue   = "\033[34m"
	purple = "\033[35m"
	cyan   = "\033[36m"
	gray   = "\033[37m"
	white  = "\033[97m"

	ANSIC      = "Mon Jan _2 15:04:05 2006"
	UnixDate   = "Mon Jan _2 15:04:05 MST 2006"
	RFC822Z    = "02 Jan 06 15:04 -0700"
	RFC3339    = "2006-01-02T15:04:05Z07:00"
	Stamp      = "Jan 02 15:04:05"
	StampMilli = "Jan 02 15:04:05.000"
	StampMicro = "Jan 02 15:04:05.000000"
	StampNano  = "Jan _2 15:04:05.000000000"
)
