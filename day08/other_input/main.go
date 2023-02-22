package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	/*
		line ，从stdin中读取一行数据（字节集合--转换成为字符串）
		reader 默认一次能4096个字节（汉字4096/3个）
			1，一次性读完，isPrefix=false
			2，先读一部分，isPrefix=true，再去读取isPrefix=false
	*/
	line, _, _ := reader.ReadLine()
	data := string(line)
	fmt.Println(data)

}
