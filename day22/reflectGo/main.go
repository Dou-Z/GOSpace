package main

import (
	"fmt"
	"reflect"
)

func reflect_examplr(a interface{}) {
	t := reflect.TypeOf(a)
	fmt.Printf("type of a is:%v\n", t)
	v := reflect.ValueOf(a)
	fmt.Printf("value of a is:%v\n", v)
	k := t.Kind()
	switch k {
	case reflect.Int64:
		fmt.Println("a is int64")
	case reflect.String:
		fmt.Println("a is string")
	case reflect.Float32:
		fmt.Println("a is Float32")
	}
}

func reflect_value(a interface{}) {
	v := reflect.ValueOf(a)
	fmt.Printf("value of a is:%v\n", v)
	k := v.Kind()
	switch k {
	case reflect.Int64:
		fmt.Println("a is int64,store value is:", v)
	case reflect.String:
		fmt.Println("a is string store value is:", v)
	case reflect.Float32:
		fmt.Println("a is Float32 store value is:", v)
	}
}

func main() {
	var x int64 = 3
	reflect_examplr(x)
	var q string = "golang"
	reflect_value(q)
}
