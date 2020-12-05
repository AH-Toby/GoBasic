package main

import "fmt"

func f1() int {
	x := 5
	defer func() {
		x++ // 2.defer，x=6，返回值=5  修改的是x不是返回值
	}()
	return x // 1.返回值赋值=x=5
}

// 5

func f2() (x int) {
	defer func() {
		x++ // 2.defer，x=6，返回值x=6
	}()
	return 5 // 1.返回值x=5
}

// 6

func f3() (y int) {
	x := 5
	defer func() {
		x++ // 2.defer，x=6，返回值y=5
	}()
	return x // 返回值y=x=5
}

// 5

func f4() (x int) {
	defer func(x int) {
		x++ // 2.defer，改变是函数中x的副本  // x=6，返回值x=5
	}(x)
	return 5 // 返回值x=5
}

// 5

// 传一个x的指针
func f5() (x int) {
	defer func(x *int) {
		(*x)++ // 2.defer，改变是函数中x的副本  // x=6，返回值x=5
	}(&x)    // &：取地址值
	return 5 // 返回值x=5
}

// 5

func main() {
	fmt.Println(f1()) // 5
	fmt.Println(f2()) // 6
	fmt.Println(f3()) // 5
	fmt.Println(f4()) // 5
	fmt.Println(f5()) // 6
}
