package main

import "fmt"

func main() {
	//回车符
	fmt.Printf("\r这是回车符")
	//换行符
	fmt.Printf("\n这是换行符")
	//制表符
	fmt.Printf("\t这是制表符")
	fmt.Println()
	//反斜杠
	s1 := "c:\\Code\\lesson1\\go.exe"
	fmt.Printf("s1:%s", s1)
	fmt.Println()
	// 双引号
	s2 := "\"c:\\Code\\lesson1\\'go.exe'\""
	fmt.Printf("s2:%s", s2)
}
