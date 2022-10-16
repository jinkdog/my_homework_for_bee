package main

import "fmt"

//自定义类型和类型别名

// type后面跟类型
type myint int     //自定义类型
type yourint = int //类型别名

func main() {
	var n myint
	n = 100
	fmt.Println(n)
	fmt.Printf("%T\n", n)
	var m yourint
	m = 200
	fmt.Println(m)
	fmt.Printf("%T\n", m)
	var c rune //类型别名，实际上就是int32，提示可以输入字符
	c = '中'    //不能用双引号
	fmt.Println(c)
	fmt.Printf("%T\n", c)
}
