package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	copyFile("target_2.txt", "target1.txt")
	fmt.Println("Copy done!")
}

func copyFile(dsName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		fmt.Println("打开源文件失败！！！")
	}
	defer src.Close()
	dst, err := os.OpenFile(dsName, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	defer dst.Close()
	res, errC := io.Copy(dst, src)
	if errC != nil {
		fmt.Println("文件复制失败，异常信息：", errC)
	} else {
		fmt.Println(res)
	}
	return
}
