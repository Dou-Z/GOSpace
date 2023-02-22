package main

import (
	"fmt"
	"runtime"
	"time"
)

var i int = 1

func calc() {
	for {
		i++
		fmt.Println(i)
	}
}
func main() {
	cpu := runtime.NumCPU()
	fmt.Println("CPU:", cpu)
	runtime.GOMAXPROCS(1)
	for i := 0; i < 10; i++ {
		go calc()
	}
	time.Sleep(time.Hour)
}
