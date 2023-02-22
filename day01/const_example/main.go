package main

import "fmt"


const(
	A=iota
	b
	c
	d=8
	e
	f=iota
	g
)
const(
	A1=iota
	A2
)
func main()  {
	fmt.Println(A)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)

	fmt.Println("A1A2")
	fmt.Println(A1)
	fmt.Println(A2)

}