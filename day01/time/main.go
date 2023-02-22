package main

import (
	"fmt"
	"time"
)

func testtime() {
	now := time.Now()
	fmt.Println(now)

	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	send := now.Second()

	fmt.Printf("%02d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, send)
}

func testTimestamp(timestamp int64) {
	timeobj := time.Unix(timestamp, 0)
	year := timeobj.Year()
	month := timeobj.Month()
	day := timeobj.Day()
	hour := timeobj.Hour()
	minute := timeobj.Minute()
	send := timeobj.Second()
	fmt.Printf("%02d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, send)
}

func testTicker()  {
	ticker:=time.Tick(2*time.Second)
	for i := range ticker{
		fmt.Printf("%v\n",i)
	}
}
func testConst()  {
	fmt.Printf("nano second:%d\n",time.Nanosecond) //纳秒
	fmt.Printf("micro second:%d\n",time.Microsecond) //微秒
	fmt.Printf("mill second:%d\n",time.Millisecond) //毫秒
	fmt.Printf(" second:%d\n",time.Second) //秒

}

func main() {
	// testtime()
	// timestamp := time.Now().Unix()
	// testTimestamp(timestamp)
	// testTicker()
	testConst()
}
