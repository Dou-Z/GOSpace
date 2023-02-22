package main

import (
	"fmt"
	"time"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
func prodyceShu(c chan int) {
	var i int = 1
	for {
		i++
		result := isPrime(i)
		if result {
			c <- i
		}
		time.Sleep(time.Second)
	}
}

func consumSushu(c chan int) {
	for v := range c {
		fmt.Printf("%d is prime \n",v)
	}
}

func main() {
	var intChan chan int = make(chan int, 10000)
	fmt.Printf("chan is %v\n",intChan)
	go prodyceShu(intChan)
	go consumSushu(intChan)

	time.Sleep(time.Hour)
}