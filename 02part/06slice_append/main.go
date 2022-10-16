package main

import "fmt"

//append()为切片追加元素

func main() {
	//s1=string[]{"a","b","c"}string位置和括号位置弄反
	//s1=[]string{"a","b","c"}应为冒号等于
	s1 := []string{"a", "b", "c"}
	fmt.Printf("s1=%v,len(s1)=%d,cap(s1)=%d\n", s1, len(s1), cap(s1))
	fmt.Println(s1, len(s1), cap(s1))

	//调用append函数必须用原来的切片量接收返回值
	//append追加元素，原来的底层数组放不下的时候，Go底层就会把底层的数组换一个
	//必须用变量接受append返回值
	s1 = append(s1, "d")
	fmt.Printf("s1=%v,len(s1)=%d,cap(s1)=%d\n", s1, len(s1), cap(s1))
	fmt.Println(s1, len(s1), cap(s1))
	s2 := []string{"e", "f", "g"}
	//s1=append{s1,s5...} forget
	//s1=append(s1,s2,...) 多了逗号
	s1 = append(s1, s2...) //...表示拆开
	fmt.Printf("s1=%v,len(s1)=%d,cap(s1)=%d\n", s1, len(s1), cap(s1))
	
}
