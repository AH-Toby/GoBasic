package main

import "fmt"

// 结构体是值类型
type person2 struct {
	name   string
	gender rune
	age    int
}

func main() {

	var p1 = &person2{}
	p1.name = "李四"
	// 语法糖
	p1.gender = '女'
	p1.age = 18
	fmt.Println(*p1)  // 取值
}
