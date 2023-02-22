package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// 只读的方式打开
	file, err := os.Open("./file.go")
	if err != nil {
		fmt.Println("open file failed err:", err)
		return

	}
	defer file.Close()
	var content []byte
	var buf [128]byte
	for {
		n, err := file.Read(buf[:])
		if err == io.EOF {
			break
		} else {
			fmt.Println(err)
		}
		content = append(content, buf[:n]...)
	}
	fmt.Println(string(content))
}
