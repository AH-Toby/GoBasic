package main

import "fmt"

func adder1(x int) func(int) int {
	return func(y int) int {
		x += y
		return x
	}
}
func main() {
	var f = adder1(10)
	fmt.Println(f(10)) //20
	fmt.Println(f(20)) //40
	fmt.Println(f(30)) //70

	f1 := adder1(20)
	fmt.Println(f1(40)) //60
	fmt.Println(f1(50)) //110
}
