package main

import "fmt"

type person1 struct {
	name string
	age int
	gender string
}

func newPerson(name string,age int,gender string) *person1 {
	return &person1{
		name:name,
		age:age,
		gender: gender,
	}
}

func main() {
	p := newPerson("张三",18,"男")
	fmt.Printf("%v",*p)
	fmt.Println(*p)
}