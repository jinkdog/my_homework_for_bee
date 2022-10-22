package main

import (
	"fmt"
	"sync"
)

//1 启动一个goroutine生成100个数到ch1中
//2 启动一个goroutine，ch1中取值计算其平方储存到ch2中
//3 再main函数中把ch2取值打印出来

var wg sync.WaitGroup

func f1(ch1 chan int) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		ch1 <- i
	}
	close(ch1)
}

func f2(ch1 <-chan int, ch2 chan<- int) {
	defer wg.Done()
	//for x := range ch1 { //要传入ch1和ch2两个通道，否则就会显示未解析
	//	ch2 <- x * x
	//}不知道为何这个部分出错
	for {
		x, ok := <-ch1
		if !ok {
			break
		}
		ch2 <- x * x
	}
	close(ch2)
}

func main() {
	a := make(chan int, 100) //	a:=make(ch1 chan int,100)声明通道变量时不用加名称
	b := make(chan int, 100)
	wg.Add(2)
	go f1(a)
	go f2(a, b)
	wg.Wait()
	for ret := range b {
		fmt.Println(ret)
	}
}
