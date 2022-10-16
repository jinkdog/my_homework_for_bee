package main

import "fmt"

// 函数的定义（有参数有返回值）
func sum(x int, y int) (ret int) {
	return x + y

}

// 没有返回值//通过再函数内内置打印函数来显示返回值
func f1(x int, y int) {
	fmt.Println(x + y)
}

// 没有参数没有返回值
func f2() {
	fmt.Println()
}

// 没有参数有返回值
func f3() int {
	return 1
}

// 参数可以命名也可以不命名
// 命名的返回值就相当于再函数中声明一个变量
// 没有命名的必须返回一个值，命名后的返回值可以不写
func f4(x int, y int) (ret int) {
	ret = x + y
	return //return ret 提前声明过ret后返回值可不加ret
}

// 多个返回值
func f5() (int, int) {
	return 1, 2
}

// 参数的类型简写：
// 当参数中连续多个参数类型一致时，可以将前面多个参数的类型省略
func f6(x, y, z int, m, n string, i, j bool) int {
	return x + y
}

// 可变长参数
// 可变长参数必须放在函数参数的最后
func f7(x int, y ...int) { //y可输入多个值
	fmt.Println(x)
	fmt.Println(y) //y的类型是slice
}

// 有多个返回值的函数
// 统一不命名可以，但不能一个命名一个不命名
func f8(x int, y int) (sum int, sub int) {
	sum = x + y
	sub = x - y
	fmt.Println(x + y)
	fmt.Println(x - y)
	//忘记了返回值
	return
}

// 不能在命名的函数中不能再声明命名函数
// GO语言中没有默认参数概念
func main() {
	//a1, b1 := sum(1, 2)错误的返回值只有一个，应用一个变量来接受
	a1 := sum(5, 6)
	fmt.Println(a1)
	f1(10, 20)
	f2()
	_, q := f5()
	fmt.Println(q)
	we, re := f5()
	fmt.Println(we, re)
	f7(1, 1234, 456, 678)
	f7(1)
	f8(1, 2)
}
