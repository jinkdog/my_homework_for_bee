package main

import "fmt"

func main() {
	//练习1 求数组[1，3，5，7，8]所有元素的和
	a1 := [...]int{1, 3, 5, 7, 8}
	sum := 0
	for _, v1 := range a1 {
		sum = sum + v1
	}
	fmt.Println(sum)

	//练习2
	//找出和为8的两个数下标分别为（0，3）（1，2）
	//定义两个for循环，外层从第一个开始遍历
	//内层for循环从外层后面开始
	//它们两个数和为8

	for i := 0; i < len(a1); i++ {
		for j := i + 1; j < len(a1); j++ {
			if a1[i]+a1[j] == 8 {
				fmt.Printf("%d,%d\n", i, j)
			}

		}
	}
}
