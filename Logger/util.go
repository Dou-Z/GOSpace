package logger

import (
	"path/filepath"
	"runtime"
)

func GetLineInfo() (fileName string, funcName string, lineNo int) {
	pc, file, line, ok := runtime.Caller(2)
	if ok {
		fName := file
		fileName = filepath.Base(fName)
		funcN := runtime.FuncForPC(pc).Name()
		funcName = filepath.Base(funcN)

		lineNo = line
	}

	return
}
