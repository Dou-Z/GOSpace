package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var x int
var wg sync.WaitGroup
var mutex sync.Mutex

func add() {

	for i := 0; i < 500; i++ {
		mutex.Lock()
		x = x + 1
		mutex.Unlock()

	}
	wg.Done()
}

var n int32

func testAtomic() {
	for i := 0; i < 500; i++ {
		// 原子操作
		atomic.AddInt32(&n, 1)
	}
	wg.Done()
}

func main() {
	start := time.Now().UnixNano()
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go testAtomic()
		// go add()
	}
	// go testAtomic()
	// go testAtomic()
	wg.Wait()
	end := time.Now().UnixNano()
	cost := (end - start) / 1000 / 1000
	fmt.Println("n:", n, "cost:", cost, "ms")
}
