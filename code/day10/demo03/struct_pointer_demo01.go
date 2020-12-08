package main

import "fmt"

// 结构体是值类型
type person1 struct {
	name   string
	gender rune
	age    int
}

func main() {

	var p1 = new(person1)
	fmt.Printf("%T\n", p1) // 查看p1的类型是结构体类型的指针
	fmt.Printf("%p\n", p1)  //查看p1内存存储的值
	fmt.Printf("%p\n", &p1)  // 查看变量p1的内存地址

	p1.name = "李四"
	// 语法糖
	p1.gender = '女'
	p1.age = 18
	fmt.Println(*p1)  // 取值
}
