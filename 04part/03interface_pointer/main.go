package main

import (
	"fmt"
)

type animal interface {
	move()
	eat(string)
}

type cat struct {
	name string
	feet int8
}

//// 使用了值接收者实现了接口的方法
//func (c cat) move() {
//	fmt.Println("猫步")
//}
//
//func (c cat) eat(food string) {
//	fmt.Printf("猫吃%s...\n", food)
//}

// 使用了指针接收者实现了接口的方法
func (c *cat) move() {
	fmt.Println("猫步")
}

func (c *cat) eat(food string) {
	fmt.Printf("猫吃%s...\n", food)
}
func main() {
	var a1 animal

	c1 := cat{"zfy", 4}
	c2 := &cat{"zfy", 4}

	//c1 = a1
	//c2 = a1

	a1 = &c1 //使用animal这个接口的是cat的指针类型
	a1 = c2

	fmt.Println((a1))
}
