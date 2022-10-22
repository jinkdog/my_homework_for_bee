package main

import (
	"fmt"
	"sync"
)

var a chan int
var wg sync.WaitGroup

func debuffchannel() {
	fmt.Println(a)
	a = make(chan int) //不带缓冲区的初始化
	wg.Add(1)
	go func() {
		defer wg.Done()
		x := <-a
		fmt.Println("后台goroutine从通道b中取出了", x)
	}()
	a <- 10 //由于无缓冲区，程序报错，所以要在之前输入
	fmt.Println("10发送到通道b当中")
	a = make(chan int, 16) //带缓冲区的初始化
	//通道必须要初始化才能使用
	fmt.Println(a)
	wg.Wait()
}

func buffchannel() {
	a = make(chan int, 16)
	a <- 10
	fmt.Println("10发送到通道a当中")
	x := <-a
	fmt.Println("后台goroutine从通道a中取出了", x)
	close(a) //关闭通道a
}

func main() {
	//fmt.Println(a)
	//a = make(chan int) //不带缓冲区的初始化
	//wg.Add(1)
	//go func() {
	//	defer wg.Done()
	//	x := <-a
	//	fmt.Println("后台goroutine从通道a中取出了", x)
	//}()
	//a <- 10 //由于无缓冲区，程序报错，所以要在之前输入
	//fmt.Println("10发送到通道a当中")
	//a = make(chan int, 16) //带缓冲区的初始化
	////通道必须要初始化才能使用
	//fmt.Println(a)
	//wg.Wait()\
	buffchannel()
}
