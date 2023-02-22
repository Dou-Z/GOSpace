package logger

const (
	LogLevelDebug = iota
	LogLevelTrace
	LogLevelInfo
	LogLevelWarn
	LogLevelError
	LogLevelFatal
)

func getLevelText(level int) string {
	switch level {
	case LogLevelDebug:
		return "DEBUG"
	case LogLevelTrace:
		return "TRACE"
	case LogLevelWarn:
		return "WARN"
	case LogLevelInfo:
		return "INFO"
	case LogLevelError:
		return "ERROR"
	case LogLevelFatal:
		return "FATAL"

	}
	return "UNKNOWN"
}
