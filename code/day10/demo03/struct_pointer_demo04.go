package main

import "fmt"

// 结构体是值类型
type person5 struct {
	name   string
	gender rune
	age    int
}

func main() {
	p1 := person5{
		name:   "张三",
		gender: '男',
		age:    18,
	}
	fmt.Printf("p1:%v\n", p1)
	p2 := &person5{
		name:   "李四",
		gender: '男',
		age:    20,
	}
	fmt.Printf("p2:%v\n", *p2)
}
