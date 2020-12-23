package main

import "fmt"

type student struct {
	name string
	age  int
}

func main() {
	m := make(map[string]*student)
	// [{name: "小王子", age: 18},{name: "娜扎", age: 23},{name: "大王八", age: 9000}]
	stus := []student{
		{name: "小王子", age: 18},
		{name: "娜扎", age: 23},
		{name: "大王八", age: 9000},
	}
	fmt.Println(stus)
	for _, stu := range stus {
		println(&stu)
		m[stu.name] = &stu
	}
	fmt.Println(m)
	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}
}