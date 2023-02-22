package main

import (
	"fmt"
)

func calc(str string) (charCount int, numCount int, spaceCount int, otherCount int) {
	utfChars := []rune(str)
	for i := 0; i < len(utfChars); i++ {
		if utfChars[i] >= 'a' && utfChars[i] <= 'z' || utfChars[i] >= 'A' && utfChars[i] <= 'Z' {
			charCount++
			continue
		}
		if utfChars[i] >= '0' && utfChars[i] <= '9' {
			numCount++
			continue
		}
		if utfChars[i] == ' ' {
			spaceCount++
			continue
		}
		otherCount++
	}
	return
}

func example3() {
	str1 := "qwert    你好 123123"
	strnum, numnum, spacenum, othernum := calc(str1)
	fmt.Printf("字符数：%d\t数字数：%d\t空格数：%d\t其他：%d\n", strnum, numnum, spacenum, othernum)
}

func main() {
	example3()
}
