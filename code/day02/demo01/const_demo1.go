package main

import "fmt"

const (
	statusOk = 200
	notFound = 404
)

func main() {
	//pi=123
	fmt.Printf("状态码：%d", statusOk)
	fmt.Println()
	fmt.Printf("状态码：%d", notFound)
}
