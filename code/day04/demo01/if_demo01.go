package main

import "fmt"

func main() {
	// if条件判断
	age := 19
	if age > 40 {
		fmt.Println("中年")
	}else if age >18 {
		fmt.Println("青年")
	}else{
		fmt.Println("少年")
	}
}
