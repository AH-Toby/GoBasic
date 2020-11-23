package main

import "fmt"

// 声明变量
/*
标准声明
var name string
var age int
var isOk bool
*/
// 批量声明
var (
	name string  // 初始值：""
	age int  // 初始值：0
	isOk bool // 初始值：false
)

func main() {
	name="Toby"
	age=18
	isOk=true

	fmt.Print(isOk)
	fmt.Println()  //快捷打印空行
	fmt.Printf("name:%s\n",name)
	fmt.Println(age)
}