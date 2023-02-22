package main

import (
	"fmt"
)

func boolIf(n int) bool {
	var (
		hundred int
		decade  int
		unit    int
	)
	hundred = n / 100
	decade = (n - hundred*100) / 10
	unit = n - hundred*100 - decade*10
	// fmt.Println(n,"====",hundred, decade, unit)
	// fmt.Println()
	if n == hundred*hundred*hundred+decade*decade*decade+unit*unit*unit {
		return true
	} else {
		return false
	}

}

func narNumber() {
	for i := 100; i < 1001; i++ {

		if boolIf(i) {
			fmt.Printf("【%d】是水仙花数 \n", i)
		}

	}
}

func main() {
	narNumber()
}
