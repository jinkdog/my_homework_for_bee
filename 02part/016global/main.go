package main

import "fmt"

var x = 100 //定义一个全局变量

//func1(){
//	fmt.Println(x)
//}函数语法格式错误

func f1() {
	//x:=10
	name := "zfy"
	//函数中查找变量的顺序
	//1 先在函数内部查找
	//2 找不到就往函数的外面查找，一直找到全局
	fmt.Println(x, name)
}
func main() {
	f1()
	//fmt.Println(name)函数内部定义的变量只能在函数内部使用
	//语句块中定义的变量只能在语句块中使用
	if i := 10; i < 10 {
		fmt.Printf("school")
	}
	//fmt.Println(i)
	for i := 10; i < 23; i++ {
		fmt.Println(i)
	}
	//fmt.Println(i)
}
