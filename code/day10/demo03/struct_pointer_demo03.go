package main

import "fmt"

// 结构体是值类型
type person4 struct {
	name   string
	gender rune
	age    int
}

func main() {
	var p1 person4
	fmt.Println(p1)
}
