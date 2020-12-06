package main

import "fmt"

type person01 struct {
	name   string
	age    int
	gender rune
	city   string
	hobby  []string
}

func main() {
	var p person01
	p.name = "张三"
	p.age = 18
	p.city = "上海"
	p.gender = '男'
	p.hobby = []string{"游戏", "电影"}
	fmt.Println(p)
	fmt.Printf("%T", p)

}
