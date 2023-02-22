package logger

import "testing"

func TestFileLogger(t *testing.T) {
	logger := NewFileLog(LogLevelDebug, "D:\\Desktop\\文件夹\\Log", "log221109")
	logger.Debug("user id[%d] is come from china", 1234)
	logger.Warn("test warn log")
	logger.Fatal("test Fatal log")
	logger.Close()

}
