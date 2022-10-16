package main

import "fmt"

// 匿名函数
func main() {
	//函数内由于无法再声明函数，于是只能利用匿名函数声明
	f1 := func(x, y int) { //f1:=func(x,y int)int其中无返回值不用再加入int
		fmt.Println(x + y)
	}
	f1(1, 2)
	//如果只调用一次的函数，还可以写成立即执行函数
	func(x, y int) {
		fmt.Println(x + y)
		fmt.Println("helloworld")
	}(100, 200) //括号表示函数立即执行，函数执行完即被销毁
}
