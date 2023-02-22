package GoLog

type logInter interface {
	LogDebug(msg string)
	LogWarn(msg string)
}
