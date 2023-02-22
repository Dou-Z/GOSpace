package calc

var a int = 100 // 小写变量外部不可访问
var A int = 200 // 大写变量外部可以访问

// 大写字母开头的函数名外部可以访问
func Add(a,b int) int{
	return a+b

}

func sub (a,b int) int{
	return a-b
}