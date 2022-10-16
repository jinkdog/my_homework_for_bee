package main

//import("fmt")格式记忆错误
import "fmt"

func main() {
	//1&地址符号
	var n = 18
	var m = 20 - n
	var k = &n
	var j = &m
	fmt.Println(k)
	fmt.Println(j)
	//2*取地址符号
	var i = *k
	var y = *j
	fmt.Println(i)
	fmt.Println(y)
	//new的使用
	//实际上使s1变成了一个地址
	//*s1实际上是一个数据，才能给它赋值
	//println和fmt.println的差别不大
	s1 := new(int)
	//s1 = 1000
	println(s1)
	*s1 = 100
	fmt.Println(s1)
	fmt.Println(*s1)
}
