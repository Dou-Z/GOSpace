package main

import (
	"errors"
	"fmt"
)

func test() {
	defer func() {
		if err := recover(); err != nil {
			println(err.(string)) // 将 interface{} 转型为具体类型。
		}
	}()
	panic("panic error!")
}

func testTry01() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	var ch chan int = make(chan int, 10)
	close(ch)
	ch <- 1

}

func testtry02() {
	defer func() {
		fmt.Println(recover())
	}()

	defer func() {
		panic("defer panic")
	}()
	panic("test panic")
}

func testTry03() {
	defer func() {
		fmt.Println(recover()) // 有效
	}()
	defer recover()              // 无效
	defer fmt.Println(recover()) // 无效
	defer func() {
		func() {
			println("defer inner")
			recover() // 无效
		}()
	}()
	// panic("test panic")
}

var ErrDivByZero = errors.New("division by zero")

func testErrorDiv(x, y int) (int, error) {
	if y == 0 {
		return 0, ErrDivByZero
	}
	return x / y, nil
}
func testError01() {
	defer func() {
		fmt.Println(recover())
	}()
	switch z, err := testErrorDiv(10, 0); err {
	case nil:
		println(z)
	case ErrDivByZero:
		panic(err)
	}
}

func main() {
	// test()
	// testTry01()
	// testtry02()
	// testTry03()
	testError01()
}
