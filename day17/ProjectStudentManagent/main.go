package main

import (
	"fmt"
	"os"
)

var studentList []*studentIfo

type studentIfo struct {
	name  string
	age   int
	sex   int
	score float32
}

func AddStudentInfo() {
	var (
		name  string
		age   int
		sex   int
		score float32
	)
	fmt.Println("请输入学生姓名：")
	fmt.Scanf("%s\n", &name)
	fmt.Println("请输入学生年龄：")
	fmt.Scanf("%d\n", &age)
	fmt.Println("请输入学生性别[0/1]：")
	fmt.Scanf("%d\n", &sex)
	fmt.Println("请输入学生考试成绩[0-100]：")
	fmt.Scanf("%f\n", &score)
	user := studentIfo{
		name:  name,
		age:   age,
		sex:   sex,
		score: score,
	}
	for index, v := range studentList {
		if v.name == user.name {
			studentList[index] = &user
			fmt.Println("=====已完成修改=====")
			return
		}
	}
	studentList = append(studentList, &user)
	fmt.Println("=====已完成添加=====")
}
func ModfiyStudent() {
	var name1 string
	var (
		age1   int
		sex1   int
		score1 float32
	)

	fmt.Println("请输入你想修改的学生姓名：")
	fmt.Scanf("%s\n", &name1)

	for _, val := range studentList {
		if val.name == name1 {
			fmt.Println("请输入修改的年龄（为0则不修改）：")
			fmt.Scanf("%d\n", &age1)
			// fmt.Println(age1)
			if age1 != 0 {
				val.age = age1
			}

			fmt.Println("请输入修改的性别（为0则不修改）：")
			fmt.Scanf("%d\n", &sex1)
			// fmt.Println(sex1)
			if sex1 != 0 {
				val.sex = sex1
			}

			fmt.Println("请输入修改的成绩（为0则不修改）：")
			fmt.Scanf("%f\n", &score1)
			// fmt.Println(score1)
			if score1 != 0 {
				val.score = score1
			}

		}
		fmt.Println("=====已完成修改=====")
	}

}
func ShowAllInfo() {
	for _, value := range studentList {
		fmt.Printf("学生信息：%v\n", value)
	}
}
func main() {
	var a int32
	for {

		fmt.Println("1.添加新学生信息")
		fmt.Println("2.修改学生信息")
		fmt.Println("3.展示所有学生信息")
		fmt.Println("4.退出")
		fmt.Println("+++++++请选择你想办事件的编号：++++++++")
		fmt.Println()
		fmt.Println()
		fmt.Scanf("%d\n", &a)
		fmt.Println(a)
		switch a {
		case 1:
			AddStudentInfo()
		case 2:
			ModfiyStudent()
		case 3:
			ShowAllInfo()
		case 4:
			os.Exit(0)
		default:
			fmt.Println("请输入正确的指令")
		}
	}

}
