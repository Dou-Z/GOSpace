package main

import (
	"fmt"
	"strings"
)

func statWordCount(str string) map[string]int {
	var result map[string]int = make(map[string]int,128)
	words := strings.Split(str, " ")
	for _, v := range words {
		count, ok := result[v]
		fmt.Println(ok)
		if !ok {
			result[v] = 1
		} else {
			result[v] = count + 1
		}

	}
	return result
}
func main() {
	str := "how do you do ? do you like she ?"
	res := statWordCount(str)
	fmt.Printf("res:%#v\n", res)
}
