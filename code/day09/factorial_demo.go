package main

import "fmt"

func factorial(n uint64)  uint64{
	if n <=1{
		return n
	}
	return n*factorial(n-1)
}
func main() {
	ret := factorial(5)
	fmt.Println(ret)
}