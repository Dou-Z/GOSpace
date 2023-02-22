package main

import "fmt"

func main() {
	// const a int = 100
	// fmt.Printf("a=%d\n", a)
	const(
		a int =100
		b
		c int = 200
		d
	)
	fmt.Printf("a=%d b=%d c=%d d=%d\n", a,b,c,d)
	/*
	const(
		e = iota
		f=iota
		g=iota
	)
	两种写法是一样的意思
	*/
	const(
		e=iota
		f
		g
	)
	fmt.Printf("e=%d f=%d g=%d\n",e,f,g)

	const(
		a1=1<<iota
		a2
		a3

	)
	fmt.Printf("a1=%d a2=%d a3=%d\n",a1,a2,a3)
}
