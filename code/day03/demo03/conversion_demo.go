package main

import "fmt"

func main() {
	a:=5.0
	b:=int(a)
	fmt.Println("a:",a)
	fmt.Printf("a:%T",a)
	fmt.Println()
	fmt.Println("b:",b)
	fmt.Printf("b:%T",b)
}
