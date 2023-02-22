package main

import (
	"flag"
	"fmt"

	"math/rand"
	"time"
)

var (
	length  int
	charset string
)

const (
	Numstr   = "0123456789"
	CharStr  = "QWERTYUIOPASDFGHJKLZXCVBNMqwertyuiopasdfghjklzxcvbnm"
	SpaceStr = " /.,<>!@#$%^&*()+=-[]{}"
)

func parseArgs() {
	flag.IntVar(&length, "l", 16, "-l 生产密码长度")
	flag.StringVar(&charset, "t", "num",
		`-t： 指定密码生成的字符集，
	num：只使用数字【0~9】，
	char：只使用英文字母【a~zA~Z】，
	mix：使用数字和字母，
	advance：使用数字、字母、以及特殊字符`)
	flag.Parse()
}
func generatePassword() string {
	var password []byte = make([]byte, length, length)
	var sourceStr string
	if charset == "num" {
		sourceStr = Numstr
	} else if charset == "char" {
		sourceStr = CharStr
	} else if charset == "mix" {
		sourceStr = fmt.Sprintf("%s%s", Numstr, CharStr)
	} else if charset == "advance" {
		sourceStr = fmt.Sprintf("%s%s%s", Numstr, CharStr, SpaceStr)
	} else {
		sourceStr = Numstr
	}
	// fmt.Println(sourceStr)
	for i := 0; i < length; i++ {
		index := rand.Intn(len(sourceStr))
		password[i] = sourceStr[index]
	}
	return string(password)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	parseArgs()
	// fmt.Printf("length:%d charset:%s\n", length, charset)
	pwd := generatePassword()
	fmt.Println("The Code is :")
	fmt.Println(pwd)
}
