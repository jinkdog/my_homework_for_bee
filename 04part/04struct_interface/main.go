package main

import "fmt"

// 同一个结构体实现多个接口
// 接口还可以嵌套
type animal interface {
	mover
	eater
}

type mover interface {
	move()
}

type eater interface {
	eat(string)
}

type cat struct {
	name string
	feet int8
}

// 实现mover接口
func (c *cat) move() {
	fmt.Println("猫步")
}

// 实现eater接口
func (c *cat) eat(food string) {
	fmt.Printf("猫吃%s...\n", food)
}
