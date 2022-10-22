package main

import "fmt"

type speaker interface {
	speak() //只要实现speak方法的类型，都是speaker类型
}

type cat struct{}

type dog struct{}

func (c cat) speak() { //cspeak 由于忘记这种函数的运算形式 导致输入错误 该问题应该与结构体有关
	fmt.Println("喵喵喵")
}

func (d dog) speak() { //同理
	fmt.Println("汪汪汪")
}

func da(x speaker) {
	x.speak()
}

func main() {
	d1 := dog{}
	c1 := cat{}
	da(d1)
	da(c1)
	var ss speaker //var ss = speaker不需要加等号
	ss = c1
	fmt.Println(ss)

}
