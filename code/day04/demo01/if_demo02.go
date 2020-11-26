package main

import "fmt"

func main() {
	// if特殊格式条件判断
	if age := 19; age > 40 {
		fmt.Println("中年")
	} else if age > 18 {
		fmt.Println("青年")
	} else {
		fmt.Println("少年")
	}
}
