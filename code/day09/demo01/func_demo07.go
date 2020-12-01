package main

import "fmt"

func sum7(x, y int) int {
	ret := x + y
	return ret
}
func main() {
	ret := sum7(1, 2)
	fmt.Println(ret)
}
