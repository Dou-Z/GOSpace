package main

import (
	"fmt"
	"time"
)

func product(c chan int) {
	c <- 100
	fmt.Println("product finished")
}
func consume(c chan int) {
	data := <-c
	fmt.Println(data)
}

func main() {
	var c chan int
	fmt.Printf("c=%v\n", c)

	c = make(chan int)
	go product(c)
	// go consume(c)
	time.Sleep(time.Hour)
}
