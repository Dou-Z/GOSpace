package main

import "fmt"

type People struct {
	name string
	city string
}

func (p People) Print() {
	fmt.Printf("name=%s city=%s\n", p.name, p.city)
}
func (p People) Set(name string, city string) {
	p.name = name
	p.city = city
	fmt.Println(p.name, p.city)
}

func (p *People) SetV2(name string, city string) {
	p.name = name
	p.city = city
}
func main() {
	var p1 People = People{
		name: "user01",
		city: "成都",
	}
	p1.Print()
	p1.Set("userz", "beijing")
	p1.Print()
	(&p1).SetV2("userz", "beijing")
	p1.Print()
}
