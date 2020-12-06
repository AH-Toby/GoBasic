package main

import "fmt"

func main() {
	var p struct{
		name,city string
		age int
		gender rune
		hobby []string
	}
	p.name = "张三"
	p.age = 18
	p.city = "上海"
	p.gender = '男'
	p.hobby = []string{"游戏", "电影"}
	fmt.Println(p)
	fmt.Printf("%T", p)
}
