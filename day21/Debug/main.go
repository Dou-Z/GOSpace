package main

import (
	"github.com/study/day21/GoLog"
)

func main() {
	/*
		file := GoLog.NewFileLog("C:/a.log")
		file.LogDebug("this is a debug")
		file.LogWarn("this is a Warn")
	*/
	// log := GoLog.NewFileLog("c:/a.txt")
	log := GoLog.NewconsoleLog("hello")
	log.LogDebug("this is a debug file")
	log.LogWarn("this is a warn file")

}
