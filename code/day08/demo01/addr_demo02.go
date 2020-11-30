package main

import "fmt"

func main() {
	var  a *int  // 声明一个int类型指针变量a
	*a =100
	fmt.Println(*a)
}
