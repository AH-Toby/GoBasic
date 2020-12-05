package main

import "fmt"

// 闭包
func f1(f func()) {
	fmt.Println("this is f1")
	f()
	fmt.Println("this is f1 end")
}
func f2(x, y int) {
	fmt.Println(x + y)
	fmt.Println("this is f2")
}

// 下面的方法就是起这个作用的
func transFunc(f func(int, int), x, y int) func() {
	return func() {
		f(x, y)
	}
}
func main() {
	// 需求要把f2传到f1中运行
	// 由于f2 和f1接收的参数类型不同所以不能够直接传入，需要通过一个方法将f2变成一个能传入的类型
	x:=100
	y:=300
	f1(transFunc(f2,x,y))  // 这样就把f2传到f1中运行了
}
