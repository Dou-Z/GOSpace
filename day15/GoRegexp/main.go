package main

import (
	"fmt"
	"regexp"
)

func testRegexp() {
	//1,字符串匹配
	m2, t := regexp.MatchString("foo.*", "seafoods")
	fmt.Println(m2, t)
	reg1 := regexp.MustCompile(`\d{11}`)
	ret1 := reg1.FindString("151215321532165541655Adawdsd123w222")

	ret2 := reg1.FindAllString("15231515210qwe4552251231532151", -1)
	fmt.Println(ret1, ret2)

	// 正则 分组
	reg2 := regexp.MustCompile(`(\d{6})(\d{4})(\d{2})(\d{2})(\d{3})([0-9]|x)`)
	res1 := reg2.FindStringSubmatch("qwe41663519980216334490xsd")
	res2 := reg2.FindAllStringSubmatch("qwe41663519980216334490xsd", -1)
	fmt.Println(res1)
	fmt.Println(res2)

}

func main() {
	testRegexp()
}
