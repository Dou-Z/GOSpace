package logger

import (
	"fmt"
	"os"
	"time"
)

type fileLog struct {
	Level    int
	FileName string
	FilePath string
	file     *os.File
	warnFile *os.File
}

func NewFileLog(level int, filePath, filename string) Loginterface {
	logger := &fileLog{
		Level:    level,
		FilePath: filePath,
		FileName: filename,
	}
	logger.init()
	return logger
}

func (f *fileLog) init() {
	filename := fmt.Sprintf("%s/%s.log", f.FilePath, f.FileName)
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)

	if err != nil {
		panic(fmt.Sprintf("open failed %s failed err:%v", filename, err))
	}
	f.file = file
	//写错误日志和fatal日志的文件
	filename = fmt.Sprintf("%s/Err_%s.log", f.FilePath, f.FileName)
	file, err = os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		panic(fmt.Sprintf("open failed %s failed err:%v", filename, err))
	}
	f.warnFile = file
}

func (f *fileLog) SetLevel(level int) {
	if level < LogLevelDebug || level > LogLevelFatal {
		level = LogLevelDebug
	}
	f.Level = level

}
func (f *fileLog) WriteLog(file *os.File, level int, format string, args ...interface{}) {
	if f.Level > LogLevelDebug {
		return
	}
	now := time.Now()
	nowStr := now.Format("2006-01-02 15:04:05.999")
	levelStr := getLevelText(level)
	fileName, funcName, lineNo := GetLineInfo()
	// fmt.Fprint(f.file, nowStr, levelStr, fileName, funcName, lineNo)
	msg := fmt.Sprintf(format, args...)

	fmt.Fprintf(file, "%s [%s] (%s %s:%d) %s\n", nowStr, levelStr, fileName, funcName, lineNo, msg)
}

func (f *fileLog) Debug(format string, args ...interface{}) {
	f.WriteLog(f.file, LogLevelDebug, format, args...)

}
func (f *fileLog) Trace(format string, args ...interface{}) {
	f.WriteLog(f.file, LogLevelTrace, format, args...)
}

func (f *fileLog) Info(format string, args ...interface{}) {
	f.WriteLog(f.file, LogLevelInfo, format, args...)
}

func (f *fileLog) Warn(format string, args ...interface{}) {
	f.WriteLog(f.warnFile, LogLevelWarn, format, args...)
}

func (f *fileLog) Error(format string, args ...interface{}) {
	f.WriteLog(f.warnFile, LogLevelError, format, args...)
}

func (f *fileLog) Fatal(format string, args ...interface{}) {
	f.WriteLog(f.warnFile, LogLevelFatal, format, args...)
}
func (f *fileLog) Close() {
	f.file.Close()
	f.warnFile.Close()
}
