package main

import "fmt"

// 结构体是值类型
type person6 struct {
	name   string
	gender rune
	age    int
}

func main() {
	p1 := person6{
		"张三",
		'男',
		18,
	}
	fmt.Printf("p1:%v\n", p1)
	p2 := &person6{
		"李四",
		'男',
		20,
	}
	fmt.Printf("p2:%v\n", *p2)
}
