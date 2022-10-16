package main

import "fmt"

//make()函数创造切片
//切片是比较类型

func main() {
	s1 := make([]int, 5, 10)
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d", s1, len(s1), cap(s1))

	s2 := make([]int, 0, 10)
	fmt.Printf("s2=%v len(s2)=%d cap(s2)=%d", s2, len(s2), cap(s2))
	//切片的赋值
	s3 := []int{1, 2, 3}
	s4 := s3
	fmt.Println(s3, s4)
	s3[0] = 1000
	fmt.Println(s3, s4)
	//切片的遍历
	//1索引遍历
	for i := 0; i < len(s3); i++ {
		fmt.Println(s3[i])
	}
	//2range遍历
	for i, v := range s3 {
		fmt.Println(i, v)
	}

}
