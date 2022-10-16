package main

import "fmt"

// 结构体//type用于定义类型//结构体是值类型
type person struct {
	name  string
	age   int
	hobby []string //切片，未初始化，不分配内存
}

func main() {
	//声明一个person类型的变量p
	var p person
	//通过字段赋值
	p.name = "zfy"
	p.age = 18
	//p.hobby=["football","badminton"]
	//数组的使用错误
	p.hobby = []string{"football", "badminton", "rap"}
	fmt.Println(p)
	//访问变量p的字段
	fmt.Println(p.name)
	var a1 person
	a1.name = "zaq"
	fmt.Println(a1.name)
	//匿名结构体
	var s struct {
		x string
		y int
	}
	s.x = "zsx"
	s.y = 9008
	fmt.Println(s)
}
