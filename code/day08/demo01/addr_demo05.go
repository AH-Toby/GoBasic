package main

import "fmt"

func main() {
	a := make([]int, 1, 10)
	a[0] = 1
	fmt.Println(a)
}
