package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type Person struct {
	Name string
	Sex  string
	Age  int
}

// var User map[string]interface{}

func login(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("../../../webhtml/index.html")
	if err != nil {
		fmt.Fprintf(w, "load login.html failed")
		return
	}
	// p := Person{
	// 	Name: "douz",
	// 	Sex:  "man",
	// 	Age:  20,
	// }
	User := make(map[string]interface{})
	User["name"] = "Mary"
	User["sex"] = "女"
	User["age"] = 18

	// t.Execute(w, p)
	t.Execute(w, User)
	t.Execute(os.Stdout, User)

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
