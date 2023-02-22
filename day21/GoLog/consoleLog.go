package GoLog

import "fmt"

type consoleLog struct{}

func NewconsoleLog(msg string) logInter {
	// fmt.Println(&consoleLog{})
	return &consoleLog{}
}

func (f *consoleLog) LogDebug(msg string) {
	fmt.Println("consoleLog", msg)
}

func (f *consoleLog) LogWarn(msg string) {
	fmt.Println("consoleLog", msg)
}
