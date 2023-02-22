package main

import "fmt"

type Animal interface {
	Talk()
	Eat()
	Name() string
}
type Dog struct {
}

func (d Dog) Talk() {
	fmt.Println("狗叫")
}
func (d Dog) Eat() {
	fmt.Printf("啃骨头\n")
}

func (d Dog) Name() string {
	fmt.Println("我叫旺财")
	return "旺财"
}

func main() {
	var a Dog
	var b Animal
	b = a

	b.Eat()

}
