package main

import (
	"fmt"
	"sort"
)

func main(){
	var a[5]int = [5]int{5,4,6,3,2}
	sort.Ints(a[:])
	fmt.Println("a:",a)

	var b [5]string = [5]string{"ac","cc","kn","ic","as"}
	sort.Strings(b[:])
	fmt.Println("b:",b)
}

