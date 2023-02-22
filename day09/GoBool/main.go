package main

import (
	"fmt"
	"strconv"
)

func main() {
	//
	v1, err := strconv.ParseBool("t")
	fmt.Println(v1, err)

	v2 := strconv.FormatBool(false)
	fmt.Println(v2)
}
