package main

import (
	"fmt"
	"time"
)

func calc() {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println("run", i, "times")
	}
}
func main() {
	go calc()
	fmt.Println("i exited")
	time.Sleep(10 * time.Second)
}
