# 第一个go语言代码

## 1.go语言的文件的基本结构

```go
package main  // 代表函数入口
//如果这个是main包那么就可以直接编译成一个可执行的二进制文件，其他做导包用

// 导入语言
import "fmt"
// 函数外只能放置标识符（变量、常量、函数、类型）的声明


// 程序入口函数
func main() {
	fmt.Println("hello world!")
}
```



## 2.标识符与关键字

### 标识符

在编程语言中具有特殊意义的词，如变量名、常量名、函数名等等。

**Go语言中的标识符:**

> 由字母，数字和`_`(下划线组成)，并且只能以字母和`_`开头。

例如：`abc`,`_`,`_123`,`a123`

### 关键字

是指在编程语言中预先定义好的具有特殊含义的标识符。关键字和保留字都不建议用作变量名。

Go语言中的25个关键词

```go
break        default      func         interface    select
case         defer        go           map          struct
chan         else         goto         package      switch
const        fallthrough  if           range        type
continue     for          import       return       var
```

此外，GO语言还有37个保留字

```go
Constants:    true  false  iota  nil

        Types:    int  int8  int16  int32  int64  
                  uint  uint8  uint16  uint32  uint64  uintptr
                  float32  float64  complex128  complex64
                  bool  byte  rune  string  error

    Functions:   make  len  cap  new  append  copy  close  delete
                 complex  real  imag
                 panic  recover
```



## 3.三个Print

> **Print:**
>
> 在终端中输出要打印的内容 
>
> **Printf：**
>
> %s：占位符，格式化打印
>
> **Println：**
>
> 打印完指定的内容后面加上一个换行符





## 4.推荐使用驼峰式命名

**推荐**：小驼峰：studentName 

 

