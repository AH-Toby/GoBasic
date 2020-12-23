package main

import "fmt"

type person struct {
	name string
	age int
	gender string
}

func main() {
	p := person{
		name:"张三",
		age:18,
		gender: "男",
	}
	fmt.Printf("%v",p)
	fmt.Println(p)
}