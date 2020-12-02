package main

import "fmt"

func f3(x,y int ) int{
	sum := x+y
	return sum
}
func main() {
	a := f3
	fmt.Printf("%T",a)  // 获取类型
}
