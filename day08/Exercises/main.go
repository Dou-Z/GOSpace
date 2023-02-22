package main

import "fmt"

func question1(){
	const num = 66
	var n int
	var flag bool
	flag = true

	for flag {
		fmt.Print("请输入你猜的数字：")
		fmt.Scan(&n)
		if n == num{
			fmt.Println("恭喜你，猜对了")
			flag = false
		}else if n < num {
			fmt.Println("小了，再试一次吧")
		}else if n > num {
			fmt.Println("大了，再试一次呢")
		}
		
	}
}
func question2(){
	for i := 1;i<11;i++{
		if i != 7{
			fmt.Println(i)
		}
	}

}
func question3(){
	for i:=1;i<=100;i++{
		if i % 2 != 0{
			fmt.Printf("奇数：%d\n",i)
		}
	}
}

func question4(){
	var sum int

	for i:=1;i<=100;i++{
		sum = sum + i
	}
	fmt.Println(sum)
}

func main() {
	// question2()
	question4()

}