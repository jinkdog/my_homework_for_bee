package main

import "fmt"

type animal interface { //接口保存分为两个部分：1值的类型2值本身
	move()
	eat(string) //这个eat和cat、chicken中的eat不同一个含餐，一个不含餐
}

type cat struct {
	name string
	feet int8
}

func (c cat) move() {
	fmt.Println("猫步")
}

func (c cat) eat(food string) {
	fmt.Println("猫粮", food)
}

type chicken struct {
	feet int8
}

func (c chicken) move() {
	fmt.Println("跳")
}

func (c chicken) eat() {
	fmt.Println("饲料")
}

func main() {
	var a1 animal //定义一个接口类型的变量
	bc := cat{    //定义一个cat类型的变量bc
		name: "zfy",
		feet: 4,
	}
	a1 = bc
	a1.eat("y=")
	fmt.Println(a1)
	fmt.Printf("%T", bc)

}
