package main

import "fmt"

type x struct {
	a int
	b int
	c int
}

func main() {
	//应该声明一个变量来接受结构体
	//结构体内部元素用逗号分隔开
	m := x{
		a: 10,
		b: 10,
		c: 10,
	}
	//地址得是接受变量的
	fmt.Printf("%p\n", &(m.a))
	fmt.Printf("%p\n", &(m.b))
	fmt.Printf("%p\n", &(m.c))
}
