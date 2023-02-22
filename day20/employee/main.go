package main

import "fmt"

type Employer interface {
	CalcSalary() float32
}
type Programer struct {
	name  string
	base  float32
	extra float32
}

func NewProframer(name string, base float32, extra float32) Programer {
	return Programer{
		name:  name,
		base:  base,
		extra: extra,
	}
}
func (p Programer) CalcSalary() float32 {
	return p.base
}
func NewSale(name string, base float32, extra float32) Sale {
	return Sale{
		name:  name,
		base:  base,
		extra: extra,
	}

}

type Sale struct {
	name  string
	base  float32
	extra float32
}

func (p Sale) CalcSalary() float32 {
	return p.base + p.extra*p.base*0.5
}

func main() {
	p1 := NewProframer("搬砖1", 1500.0, 0)
	p2 := NewProframer("搬砖2", 1600.0, 0)
	p3 := NewProframer("搬砖3", 1650.0, 0)

	s1 := NewSale("销售1", 800, 2.5)
	s2 := NewSale("销售2", 700, 2.5)
	s3 := NewSale("销售3", 850, 2.5)
	fmt.Println("员工工资为：")
	fmt.Println(p1.CalcSalary())
	fmt.Println(p2.CalcSalary())
	fmt.Println(p3.CalcSalary())
	fmt.Println(s1.CalcSalary())
	fmt.Println(s2.CalcSalary())
	fmt.Println(s3.CalcSalary())

}
