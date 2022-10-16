package main

import "fmt"

//闭包是一个函数，这个函数包含了它外部作用域的一个变量

// 底层原理
// 1.函数可以作为返回值
// 2.函数内部查找变量的顺序，先在自己内部找，找不到就在外部找
func addre() func(int) int { //unc addre func(int)int{ 忘记给addre函数加括号来表示无参数输入
	var x = 100
	return func(y int) int {
		x += y //x=x+y
		return x
	} //返回值是函数所以变量接受后可以调用
}

func addre1(x int) func(int) int { //func addre1() func(x int) int {增加x为参数而不是返回值
	return func(y int) int {
		x += y //x=x+y
		return x
	}
}

// f1调用函数的函数
func f1(f func()) {
	fmt.Println("f1")
	f()
}

// f2起实际作用的函数
func f2(x, y int) { //func f2(x,y int)int不需要返回值int只需要操作函数即可
	fmt.Println("f2")
	fmt.Println(x + y)
}

// f3转换f2返回给f1
func f3(f func(int, int), x, y int) func() {
	temp := func() {
		f(x, y)
	}
	return temp
}
func main() {
	ret := addre1(100)
	ret2 := ret(200)
	fmt.Println(ret2)

	ret3 := f3(f2, 100, 200) //把原来需要传递两个int类型的参数包装成一个不需要传参的函数
	f1(ret3)
}
