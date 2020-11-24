package main

import "fmt"

const (
	c1 = iota  // 0
	c2  // 1
	_  // 2
	c3  // 3
)

func main() {
	fmt.Println("c1:", c1)
	fmt.Println("c2:", c2)
	fmt.Println("c3:", c3)
}
