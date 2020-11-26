package main

import "fmt"

func main() {
	age := 18
	switch {
	case age > 20:
		fmt.Println("需要上班")
	case age < 70 && age > 20:
		fmt.Println("需要上班")
	default:
		fmt.Println("爽呆呆")
	}
}
