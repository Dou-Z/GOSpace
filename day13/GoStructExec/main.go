package main

import "fmt"

type School struct {
	band string
	city string
}

func main() {
	var schooList []School
	for {
		var bands, city string
		fmt.Printf("请输入品牌名：")
		fmt.Scanf("%s", &bands)
		if bands == "Q" {
			break

		}
		fmt.Printf("请输入城市：")
		fmt.Scanf("%s", &city)
		fmt.Println(city)
		sch := School{band: bands, city: city}
		schooList = append(schooList, sch)

	}
	fmt.Println(schooList)
}
