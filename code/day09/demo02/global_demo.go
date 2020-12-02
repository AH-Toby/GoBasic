package main

import "fmt"

var x = 100 //定义一个全局变量
func f0() {
	//x := 10
	//函数中查找变量的顺序
	// 1.先在函数内部查找
	// 2.找不到在往函数外面查找，一直找到全局
	fmt.Println(x)
}
func main() {
	f0()
}
