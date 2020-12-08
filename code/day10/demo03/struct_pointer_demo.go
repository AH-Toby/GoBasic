package main

import "fmt"

// 结构体是值类型
type person struct {
	name string
	gender rune
	age int
}

func f2(p person)  {
	p.name = "李四"
}

func f3(p *person){
	//(*p).name = "李四"  //根据指针找到那个原变量，修改的就是原来的变量
	p.name = "李四"  // 语法糖，自动根据指正找对应的变量
}
func main() {
	var p1 person
	p1.name = "张三"
	p1.gender = '男'
	p1.age = 18
	fmt.Println(p1)
	// 需求在f2函数中修改name的值
	f2(p1)
	fmt.Println(p1)

	f3(&p1)
	fmt.Println(p1)
}


