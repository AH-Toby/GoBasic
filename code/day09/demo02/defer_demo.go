package main

import (
	"fmt"
)

func deferDemo(x,y int ) (sum int){
	fmt.Println("函数开始")
	defer fmt.Println("defer运行")
	sum = x + y
	fmt.Println("函数结束")
	return
}

func main() {
	println(deferDemo(10,20))
}