package main

import "fmt"

//defer多用于文件释放资源

func deferdemo() {
	fmt.Println("a")
	defer fmt.Println("b") //一个函数可以有多个defer语句
	defer fmt.Println("c") //多个defer语句按照先进后出执行
	fmt.Println("d")
	//无返回值，无参数
}

func main() {
	deferdemo()

}
