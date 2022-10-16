package main

import "fmt"

func f1() {
	fmt.Println()
}

func f2() int {
	return 10
}

// 函数也可以作为参数的类型
func f3(x func() int) {
	//ret:=func() func是类型，返回值是x()
	ret := x()
	fmt.Println(ret)
}

// 函数还可以作为返回值
func f4(x func() int) func(int, int) int {
	//ret:=ff应返回函数
	return ff
}

func ff(a, b int) int {
	return a + b
}
func main() {
	a := f1
	fmt.Printf("%T\n", a)
	b := f2
	fmt.Printf("%T\n", b)
	f3(f2)
	f3(b)
	f5 := f4(f2)
	fmt.Printf("%T\n", f5)
}
