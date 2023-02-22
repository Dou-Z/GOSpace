package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Person struct {
	Name string
	Sex  string
	Age  int
}

func CreateJsonData(filename string) (err error) {
	var persons []*Person
	for i := 0; i < 10; i++ {
		p := &Person{
			Name: fmt.Sprintf("name%d", i),
			Sex:  "women",
			Age:  i,
		}
		persons = append(persons, p)
	}
	data, err := json.Marshal(persons)
	if err != nil {
		fmt.Printf("=maeshal failed ,err:%s\n", err)
		return
	}
	err = ioutil.WriteFile(filename, data, 0755)
	if err != nil {
		return
	}
	return
}
func readJson(filename string) (err error) {
	var person []*Person
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &person)
	if err != nil {
		return
	}
	for _, v := range person {
		fmt.Printf("%#v\n", v)

	}
	fmt.Printf("%#v\n", person)
	return
}

func main() {
	filepath := "./person.txt"
	err := CreateJsonData(filepath)
	if err != nil {
		return
	}
	readJson(filepath)
}
