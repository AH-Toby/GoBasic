package main

import "fmt"

func f2() {
	i := 20
	if i := 10; i < 18 {
		fmt.Println("未成年")
	}
	fmt.Println(i)
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}
func main() {
	f2()
}
