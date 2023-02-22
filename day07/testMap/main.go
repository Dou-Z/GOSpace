package main

import (
	
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func testMapBase() {
	var a map[string]int
	fmt.Printf("a:%v\n", a)
	if a == nil {
		a = make(map[string]int, 20)
		a["st1"] = 100
		a["st2"] = 200
		a["st3"] = 300
		fmt.Printf("a:%v\n", a)

		b := a
		b["st1"] = 1
		fmt.Printf("after modify a : %v\n", a)
	}

}

func testsortMap() {
	rand.Seed(time.Now().UnixNano())
	var a map[string]int = make(map[string]int, 1024)

	for i := 0; i < 128; i++ {
		key := fmt.Sprintf("stu%d", i)
		value := rand.Intn(1000)
		a[key] = value
	}
	var keys []string = make([]string,0, 128)
	for key, value := range a {
		fmt.Printf("map[%s]=%d\n", key, value)
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, value := range keys {
		fmt.Printf("key:%s val :%d\n", value, a[value])
	}

}

func main() {
	// testMapBase()
	testsortMap()
}
