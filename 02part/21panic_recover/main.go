package main

import "fmt"

func funcA() {
	fmt.Println("a")
}
func funcB() {

	//defer func() {
	//	fmt.Println("erroer")
	//}() //defer 中的表达式必须为函数调用
	//panic("erro") //panic语句后的不会执行，defer语句则会接着执行，但如果defer在panic后则defer也不会执行
	//fmt.Println("b")
	defer func() {
		err := recover()
		fmt.Println(err)
		fmt.Println("relife")
	}()
	panic("erro")
	fmt.Println("b")
}
func funcC() {
	fmt.Println("c")
}
func funcD() {
	fmt.Println("d")
}

func main() {
	funcA()
	funcB()
	funcC()
	funcD()
}
