package main

import "fmt"

type Anmial struct {
	Name string
	Kind string
}
type Dog struct {
	Feet string
	*Anmial
}

func (a Anmial) talk() {
	fmt.Printf("I am talk,i am a %s\n", a.Name)
}
func (d *Dog) Eat() {
	fmt.Println("dog is eat")
}

func main() {
	var d *Dog = &Dog{
		Feet: "eat",
		Anmial: &Anmial{
			Name: "dogs",
			Kind: "four leg",
		},
	}
	d.Eat()
	d.talk()
}
