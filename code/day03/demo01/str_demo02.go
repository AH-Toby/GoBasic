package main

import "fmt"

func main() {
	// 多行字符串
	s1 := `床前明月光，
疑是地上霜。
举头望明月，
低头思故乡。`
	fmt.Printf(s1)
	// 注意多行打印是原样打印，原来是啥样打印出来就是啥样子
	s2 := `"c:\\Code\\lesson1\\'go.exe'\"`
	fmt.Printf(s2)
}
