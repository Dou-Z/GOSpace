package main

import("fmt")

func testswitch()  {
	switch a :=2;a{
	case 1:
		fmt.Printf("a=1\n")
	case 2:
		fmt.Printf("a=2\n")
	case 3:
		fmt.Printf("a=3\n")
	case 4:
		fmt.Printf("a=4\n")
	case 5:
		fmt.Printf("a=5\n")
	default:
		fmt.Printf("a=%v\n",a)
	}
	
}

func main()  {
	testswitch()
}