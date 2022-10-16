package main

import "fmt"

// Go语言中推荐使用驼峰式命名
var studentName string

// 声明变量
/*var name string
var age int
var isok bool*/

// 批量声明
var (
	name string
	age  int
	isok bool
)

func main() {
	name = "zfy"
	age = 18
	isok = true
	//Go语言中声明的变量必须使用，否则无法编译
	fmt.Print(isok)            //在终端中输出指定内容
	fmt.Println()              //快捷打印出一个空行
	fmt.Printf("name%s", name) //%s为一个占位符用name这个变量去替换
	fmt.Println(age)           //打印指定位置后会加一个换行符
	//简短变量声明，只能在函数里使用，不使用就无法编译
	s3 := "hhh"
	fmt.Println(s3)
}

// 声明变量的同时赋值
var s1 string = "zfy01"

// 类型推导
var s2 = "20"
