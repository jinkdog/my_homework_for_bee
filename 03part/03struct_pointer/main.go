package main

import "fmt"

type person struct {
	name, gender string
}

// go语言中函数参数永远是拷贝
func f(x person) {
	x.gender = "女" //修改的是副本的gender
}

// 为什么不能修改为女 因为f修改的是副本
func f2(x *person) {
	//(*x).gender = "女" //根据内存地址找原变量，修改的是原来的变量
	x.gender = "女" //语法糖，由于go中的指针是无法修改的，检测到输入的为指针时自动切换变量，意义同上
}
func main() {
	var p person
	p.name = "zfy"
	p.gender = "男"
	f(p)
	fmt.Println(p.gender)
	f2(&p)
	fmt.Println(p.gender)
	//结构体指针1
	var p2 = new(person)
	p2.name = "zxc"
	//(*p2).name = "zxc"
	fmt.Printf("%T\n", p2)
	fmt.Printf("%x\n", p2)
	fmt.Printf("%p\n", p2)  //p2保存的值是一个内存地址
	fmt.Printf("%p\n", &p2) //求p2的内存地址
	//结构体指针2
	//使用key——value初始化
	var p3 = person{
		//name=""语法错误
		name:   "sd", //别忘记逗号
		gender: "女",
	}
	fmt.Println(p3)
	//使用值列表的形式初始化，必须按定义的顺序来输入
	p4 := person{
		"zfy",
		"男",
	}
	fmt.Println(p4)

}
