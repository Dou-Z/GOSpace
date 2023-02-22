package main

import (
	"flag"
	"fmt"
)

var rec bool
var test string
var level int

func init() {
	flag.BoolVar(&rec, "r", false, "rec xxx")
	flag.StringVar(&test, "t", "default string", "string option")
	flag.IntVar(&level, "l", 1, "level of xxx")

	flag.Parse()
}

func main() {
	fmt.Printf("rec :%v\n", rec)
	fmt.Printf("test :%v\n", test)
	fmt.Printf("level :%v\n", level)
}
