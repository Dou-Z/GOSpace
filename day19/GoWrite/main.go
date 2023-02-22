package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func testWrite() {
	file, err := os.OpenFile("./test.txt", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file failed,err:", err)
	}
	defer file.Close()
	str := "hello world"
	file.Write([]byte(str))
	file.WriteString(str)
}
func testBufioWrite() {
	file, err := os.OpenFile("./test.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file failed,err:", err)
	}
	defer file.Close()
	str := "hello world\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString(str)
	}
	writer.Flush()

}

func testIoutilWrite() {
	str := "hello 成都金牛\n"
	err := ioutil.WriteFile("./test1.txt", []byte(str), 0755)
	if err != nil {
		fmt.Println("Write file failed ,err:", err)
		return
	}
}

func main() {
	// testBufioWrite()
	testIoutilWrite()

}
