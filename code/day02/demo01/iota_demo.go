package main

import "fmt"

const (
	b1 = iota  // 0
	b2  // 1
	b3  // 2
	b4  // 3
)

func main() {
	fmt.Println("b1:", b1)
	fmt.Println("b2:", b2)
	fmt.Println("b3:", b3)
	fmt.Println("b4:", b4)
}
