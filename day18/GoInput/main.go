package main

import "fmt"

func testSscan() {
	var (
		a int
		b string
		c float32
	)
	var str = "12 helloworld 33.3"
	fmt.Sscan(str, &a, &b, &c)
	fmt.Printf("a=%d,b=%s,c=%f\n", a, b, c)

}

func testSscanf() {
	var (
		a int
		b string
		c float32
	)
	var str = "99 world 66.66"
	fmt.Sscanf(str, "%d%s%f\n", &a, &b, &c)
	fmt.Printf("a=%d,b=%s,c=%f\n", a, b, c)
}
func testSscanln() {
	var (
		a int
		b string
		c float32
	)
	var str = "100 GoLang 88.8"
	fmt.Sscanln(str, &a, &b, &c)
	fmt.Printf("a=%d,b=%s,c=%f\n", a, b, c)

}

func main() {
	testSscan()
	testSscanf()
	testSscanln()
}
