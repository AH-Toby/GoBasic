package main

import "fmt"
// 相当于在函数中申明一个变量ret
func sum5(x int, y int ) (ret int){
	ret = x+y
	return
}
func main() {
	fmt.Println(sum5(1,2))
}
