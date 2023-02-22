package main

import "fmt"

func sendData(sendch chan<- int) {
	sendch <- 10
}
func readData(sendch <-chan int) {
	data := <-sendch
	fmt.Println(data)
}
func main() {
	chn1 := make(chan int)
	go sendData(chn1)
	readData(chn1)
}