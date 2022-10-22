package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func f() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		r1 := rand.Int()
		r2 := rand.Intn(11) //0<=x<11
		fmt.Println(r1, r2) //负数的随机数fmt.Println(0-r1, 0-r2)
	}
}

func f1(i int) {
	defer wg.Done()
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(13)))
	fmt.Println(i)
}

var wg sync.WaitGroup

func main() {
	//f()
	//两种解决如何判断10个goroutine都走完的方法
	//方法1：wg.Add(10)
	for i := 0; i < 10; i++ {
		wg.Add(1) //方法2
		go f1(i)
	}
	wg.Wait() //等待wg的计数器减为0
}
