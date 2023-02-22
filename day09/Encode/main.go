package main

import (
	"fmt"
	"strconv"
)

func main() {

	name := "武佩奇"

	fmt.Println(name[0],strconv.FormatInt(int64(name[0]),2))
	fmt.Println(name[1],strconv.FormatInt(int64(name[1]),2))
	fmt.Println(name[2],strconv.FormatInt(int64(name[2]),2))

	fmt.Println(name[3])
	fmt.Println(name[4])
	fmt.Println(name[5])

	fmt.Println(name[6])
	fmt.Println(name[7])
	fmt.Println(name[8])
	// fmt.Println(name[0])
	// fmt.Println(name[0])

}
