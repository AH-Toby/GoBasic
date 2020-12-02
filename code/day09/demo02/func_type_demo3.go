package main

import "fmt"

func f6(x, y int) (sum int) {
	sum = x + y
	return
}
func f7(x int) int {
	ret := x + 1
	return ret
}
func fr(x, y int, z func(int, int) int) func(int) int {
	return f7
}

func main() {
	a := fr(10,20, f6)
	b := a(20)
	fmt.Println(b)
}
