package GoLog

import "fmt"

type fileLog struct{}

func NewFileLog(file string) logInter {
	return &fileLog{}
}

func (f *fileLog) LogDebug(msg string) {
	fmt.Println("file", msg)
}

func (f *fileLog) LogWarn(msg string) {
	fmt.Println("file", msg)
}
