package main

import (
	"fmt"
	"strings"
	"github.com/study/day01/access"
)

func testBool() {
	var a bool
	fmt.Println(a)
	a = true
	fmt.Println(a)
	a = !a
	fmt.Println(a)

	var b bool = true
	if a == true && b == true {
		fmt.Println("right")
	} else {
		fmt.Println("not right")
	}

	if a == true || b == true {
		fmt.Println("|| right")
	} else {
		fmt.Println("||not right")
	}
	fmt.Printf("%t %t\n", a, b)
}
func testInt() {
	var a int8
	a = 18
	fmt.Printf("a=\n", a)
	a = -18
	fmt.Printf("a=%d\n", a)
	var b int
	b = 9999999999
	fmt.Printf("b=\n", b)


}

func testStr(){
	var a string
	a="a:hello"
	fmt.Println(a)

	b:=a
	fmt.Println(b)

	c:="c:hello"
	fmt.Println(c)

	d := `D:\nhello`
	fmt.Printf("a=%v b=%v c=%v d=%v\n",a,b,c,d)
	// 字符串拼接
	c = c+d
	fmt.Printf("c=%s\n",c)
	ips:="10.108.34.1;10.108.34.2"
	// 字符串分割
	ipArray:=strings.Split(ips,";")
	fmt.Printf("ipArray:%v\n",ipArray)
	fmt.Printf("first ip :%s\n",ipArray[0])
	fmt.Printf("second ip :%s\n",ipArray[1])
	// 字符串判断包含
	result:=strings.Contains(ips,"10.108.34.1")
	fmt.Println(result)
	// 字符串开头判断
	str1:="http://www.baidu.baidu.com"

	if strings.HasPrefix(str1,"https"){
		fmt.Printf("str is http url")
	}else{
		fmt.Println("str is not http url")
	}
	// 判断结尾
	if strings.HasSuffix(str1,"com"){
		fmt.Printf("str is com url\n")
	}else{
		fmt.Println("str is not com url\n")
	}
	// 字符出现的位置
	index:=strings.Index(str1,"baidu")
	fmt.Printf("baidu is index:%d\n",index)
	index2:= strings.LastIndex(str1,"baidu")
	fmt.Printf("baidu last index:%d\n",index2)

	// join拼接
	var strArr [] string=[]string{"10.10.0.2","10.10.0.3","10.10.0.4"}
	resultStr:=strings.Join(strArr,";")
	fmt.Printf("resultStr=%s\n",resultStr)
	fmt.Println(strArr)

}

func testAccess(){
	fmt.Printf("access.A=%d\n",access.A)
}

func main() {
	// testBool()
	// testInt()
	// testStr()
	testAccess()

}
