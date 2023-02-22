package main

import (
	"fmt"
)

func sumArr(a [10]int) int {
	var sum int = 0
	for _, val := range a {
		sum += val
	}
	return sum
}

func testArraySum() {
	var b [10]int
	for i := 0; i < len(b); i++ {
		b[i] = i
	}
	sum := sumArr(b)

	fmt.Printf("sum=%d\n", sum)
}

func TwoSum(a [5]int, target int) {
	for i := 0; i < len(a); i++ {
		other := target - a[i]
		for j := i+1; j < len(a); j++ {
			if a[j] == other {
				fmt.Printf("(%d %d)\n", i, j)
			}
		}
	}
}
func testTwoSum() {
	b := [...]int{1, 3, 5, 8, 7}
	TwoSum(b, 8)
}

func main() {
	// testArraySum()
	testTwoSum()
}
