package main

import "fmt"

func funcD() {
	fmt.Println("func D")
}

func funcE() {
	defer func() {
		err := recover()
		//如果程序出出现了panic错误,可以通过recover恢复过来
		if err != nil {
			fmt.Println("recover in E")
		}
	}()
	panic("panic in E")
}

func funcF() {
	fmt.Println("func F")
}
func main() {
	funcD()
	funcE()
	funcF()
}