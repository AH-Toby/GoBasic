package main

import "fmt"

func main() {
	day := 4
	switch day {
	case 1:
		fmt.Println("星期1")
	case 2:
		fmt.Println("星期2")
	case 3:
		fmt.Println("星期3")
	case 4:
		fmt.Println("星期4")
		fallthrough
	case 5:
		fmt.Println("星期5")
	case 6:
		fmt.Println("星期6")
	case 7:
		fmt.Println("星期7")
	default:
		fmt.Println("输入错误")
	}
}
