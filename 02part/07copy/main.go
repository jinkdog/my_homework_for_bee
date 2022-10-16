package main

import "fmt"

//copy

func main() {
	a1 := []int{1, 2, 3}
	a2 := a1
	//var a3[]int这样声明过后没有内存不能复制数据到a3
	//var a3 make([]int,0,3)忘记等于号
	var a3 = make([]int, 3, 3)
	//copy(a1, a3)长度由后面的传递到前面的
	copy(a3, a1)
	fmt.Println(a1, a2, a3)
	a1[0] = 100 //a1和a2改变，a3不变
	fmt.Println(a1, a2, a3)

	//将a1中索引为1的3这个元素删除
	a1 = append(a1[:1], a1[2:]...)
	fmt.Println(a1)
	fmt.Println(cap(a1))

	x1 := [...]int{1, 3, 5}
	//s1:=x1s1是切片而x1是数组
	s1 := x1[:]
	fmt.Println(s1, len(s1), cap(s1))
	fmt.Printf("%p/n", &s1)
	//1切片不保存具体的值
	//2切片对应一个底层数组
	//3底层数组都是占用一块连续的内存
	//s1:=append(s1[:1],s1[:2]...)冒号等于需要s1为变量？
	s1 = append(s1[:1], s1[2:]...)
	fmt.Printf("%p/n", &s1)
	fmt.Println(s1, len(s1), cap(s1))
	s1[0] = 100 //切片可以修改底层数组
	fmt.Println(x1)

}
