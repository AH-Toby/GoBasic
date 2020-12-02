package main

import "fmt"

func f4(x,y int ) int{
	sum := x+y
	return sum
}
func main() {
	var a func(int, int) int
	a = f4
	b := f4
	fmt.Printf("%T\n",a)  // 获取类型
	fmt.Printf("%T",b)  // 获取类型
}
