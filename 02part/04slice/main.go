package main

import (
	"fmt"
)

func main() {
	//切片定义
	var s1 []int    //定义存放int的切片
	var s2 []string //定义存放string的切片
	fmt.Println(s1, s2)
	fmt.Println(s1 == nil) //nil表示为空
	fmt.Println(s2 == nil)
	//切片的初始化
	s1 = []int{1, 2, 3}
	s2 = []string{"ni", "wo", "ta"}
	fmt.Println(s1, s2)
	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)
	//长度和容量
	fmt.Printf("len(s1):%d,cap(s1):%d", len(s1), cap(s1))
	fmt.Printf("len(s2):%d,cap(s2):%d", len(s2), cap(s2))

	//由数组得到切片
	a1 := [...]int{1, 2, 3, 4, 5, 6}
	s3 := a1[0:4]
	fmt.Println(s3)
	s4 := a1[1:6]
	fmt.Println(s4)
	s5 := a1[:3] //[0:3]
	s6 := a1[4:] //[4:len(a1)]
	s7 := a1[:]  //[0:len(a1)]
	fmt.Println(s5, s6, s7)
	//切片的容量指底层数组的容量
	fmt.Printf("len(s5):%d,cap(s5):%d", len(s5), cap(s5))
	//底层数组从第一个元素到最后一个元素的数量
	fmt.Printf("len(s6):%d,cap(s6):%d", len(s6), cap(s6))
	//切片再切片
	s8 := s6[1:]
	fmt.Printf("len(s8):%d,cap(s8):%d", len(s8), cap(s8))
	//切片是引用类型，都指向了一个底层数组
	fmt.Println("s6:", s6)
	a1[5] = 1300 //修改了底层数组的值
	fmt.Println("s6:", s6)
	fmt.Println("s8:", s8)

}
