package main

import "fmt"

const (
	d1 = iota  // 0
	d2 = 100  //100
	d3  //100
	d4 = iota  // 3
	d5  // 4
)

func main() {
	fmt.Println("d1:", d1)
	fmt.Println("d2:", d2)
	fmt.Println("d3:", d3)
	fmt.Println("d4:", d4)
	fmt.Println("d5:", d5)
}
