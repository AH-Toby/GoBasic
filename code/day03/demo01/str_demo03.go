package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "衣带渐宽终不悔"
	// 这里的长度是字节数
	s2 := "hello world"
	fmt.Printf("s1字符串长度：%d", len(s1))
	fmt.Println()
	fmt.Printf("s2字符串长度：%d", len(s2))
	fmt.Println()
	// 拼接字符串
	s3 := "为伊消得人憔悴"
	fmt.Printf("s1与s3利用+拼接后字符串：%s", s1+s3)
	fmt.Println()
	fmt.Printf("s1与s3利用fmt.Sprintf拼接后字符串：%s", fmt.Sprintf("%s，%s。", s1, s3))
	fmt.Println()
	// 分割
	s4 := "这 是 分 割 效 果"
	s4Split := strings.Split(s4," ")
	fmt.Println("分割后：",s4Split)
	// 判断包含
	fmt.Println(strings.Contains(s1,"终"))
	fmt.Println(strings.Contains(s2,"o"))
	// 前缀/后缀判断
	fmt.Println(strings.HasPrefix(s1,"衣"))
	fmt.Println(strings.HasSuffix(s2,"d"))
	// 子串出现的位置
	// 字节位置
	fmt.Println(strings.Index(s1,"终"))
	fmt.Println(strings.Index(s2,"o"))
	//join操作
	fmt.Println(strings.Join(s4Split,"+"))
}
