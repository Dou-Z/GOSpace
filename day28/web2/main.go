package main

import (
	"fmt"
	"html/template"
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

	t, err := template.ParseFiles("../../../webhtml/index2.html")
	if err != nil {
		fmt.Fprintf(w, "load login.html failed")
		return
	}
	p := Person{
		Name: "douz",
		Sex:  "man",
		Age:  20,
		Address: Address{
			City:     "成都",
			Province: "四川",
		},
	}
	// User := make(map[string]interface{})
	// User["name"] = "Mary"
	// User["sex"] = "女"
	// User["age"] = 18

	t.Execute(w, p)

	t.Execute(os.Stdout, p)

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
