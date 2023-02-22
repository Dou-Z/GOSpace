package main

import (
	"fmt"
	"time"
)

func hello(c chan bool) {
	time.Sleep(time.Second * 2)
	fmt.Println("hello channal")

	c <- false
}

func main() {
	var exitchan chan bool
	exitchan = make(chan bool)
	go hello(exitchan)
	fmt.Println("main thread terminate")
	<-exitchan

}
