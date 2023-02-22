package main

import (
	"fmt"
	"time"
)

func producer(chn1 chan int) {
	for i := 0; i < 10; i++ {
		chn1 <- i
		time.Sleep(time.Second)
	}
	close(chn1)
}
var li [3]int


func main() {
	fmt.Println(li)
	ch := make(chan int)
	go producer(ch)
	for v := range ch {
		fmt.Println("receive:", v)
	}
}
