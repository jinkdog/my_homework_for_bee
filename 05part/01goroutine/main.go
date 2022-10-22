package main

import (
	"fmt"
	"time"
)

func hello(i int) {
	fmt.Println("hello", i)
}

// 程序启动后会创建一个主goroutine去执行
func main() {
	for i := 1; i < 10; i++ {
		//go func() {
		//	fmt.Println(i)
		//}() //括号表示立即执行
		//调用闭包中外层的i由于for循环和goroutine执行速度不一样，所以有可能重复打印相同数字
		//解决方法
		go func(i int) {
			fmt.Println(i)
		}(i) //直接使用函数参数中的i而不是函数外的i
		//go hello(i)//开启一个单独的goroutine去执行hello函数
	}
	fmt.Println("main")
	time.Sleep(time.Second)
} //main函数结束后，由main函数启动的goroutine也都结束
//goroutine对应的函数结束了，goroutine就结束了
