package main

import (
	"fmt"
)

func deferDemo1(x,y int ) (sum int){
	fmt.Println("函数开始")
	defer fmt.Println("defer1运行")
	defer fmt.Println("defer2运行")
	defer fmt.Println("defer3运行")
	sum = x + y
	fmt.Println("函数结束")
	return
}

func main() {
	println(deferDemo1(10,20))
}