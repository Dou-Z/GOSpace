package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"os"
)

type Address struct {
	City     string
	Province string
}
type Person struct {
	Name string
	Sex  string
	Age  int
	Address
}

func login(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("./index.html")
	if err != nil {
		fmt.Fprintf(w, "load login.html failed")
		return
	}
	// p := Person{
	// 	Name: "douz",
	// 	Sex:  "man",
	// 	Age:  20,
	// 	Address: Address{
	// 		City:     "成都",
	// 		Province: "四川",
	// 	},
	// }
	var userList []*Person
	for i := 0; i < 30; i++ {
		user := Person{
			Name: fmt.Sprintf("Mary%d\n", i),
			Sex:  "women",
			Age:  rand.Intn(30),
			Address: Address{
				City:     "北京",
				Province: "北京市",
			},
		}
		userList = append(userList, &user)
	}
	fmt.Println(userList)
	err = t.Execute(w, userList)
	if err != nil {
		fmt.Printf("Execute template failed,err:%v\n", err)
	}
	t.Execute(os.Stdout, userList)

}

func main() {
	fmt.Println("index is opening!!!")
	http.HandleFunc("/index", login)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("listen server failed,err:%v\n", err)
		return
	}
}
