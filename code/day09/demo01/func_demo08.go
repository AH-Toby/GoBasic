package main

import "fmt"

func sum8(x int ,y ...int){
	fmt.Println(x)
	fmt.Println(y)
}
func main() {
	sum8(1, 2, 3, 4, 5, 6, 7)
}
