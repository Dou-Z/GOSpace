package main

import (
	"bufio"
	"fmt"
	"os"
)

func testOsStdin() {
	var buf [20]byte
	// fmt.Println(buf)
	os.Stdin.Read(buf[:]) // ==Scan输入
	// fmt.Println(buf)
	// fmt.Println(string(buf[:]))
	os.Stdout.WriteString(string(buf[:])) //= 输出
}

func testBufio() {
	var str string
	// fmt.Scan(&str)
	/*
		fmt.Scanf("%s\n", &str) //同上
		fmt.Println("read for fmt:", str)
	*/

	reader := bufio.NewReader(os.Stdin)
	str, _ = reader.ReadString('\n')
	fmt.Println("read from bufio:", str)
}
func main() {
	testBufio()

}
