package main

import "fmt"

// Go语言中函数的return不是原子操作，再底层分为两步执行
// 1 返回值赋值
// defer
// 2 真正的RET返回
// 函数中如果存在defer，那么defer执行的时机在第一步和第二步之间
func f1() int { //出现疑问：这个是返回值int还是参数int 是返回值in
	x := 5
	defer func() {
		x++
	}() //这里的括号是什么意思？ 执行函数
	return x //返回值是5 然后defer执行x=6 最后真正ret就是没有改变的返回值5
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}
func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	//return y 这样的返回值就是没有赋值过的0
	return x //1 返回值=y=x=5 2 defer修改的是x 3 真正的返回
}

func f4() (x int) {

	//defer func()(x int){这样书写就使函数已经有了参数和返回值
	defer func(x int) { //改变的是函数中x的副本
		x++
	}(x) //函数传参传入的是副本

	fmt.Println(x)
	return 5
} //不明白
func f5() (x int) { //func f6() (x *int)
	defer func(x *int) { //defer func(*int)
		(*x)++
	}(&x)
	return 5 //1 返回值x=5 2 defer x=6 3 ret返回
}
func main() {
	//a := f1()
	//fmt.Println(a)
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
	fmt.Println(f5())
}
