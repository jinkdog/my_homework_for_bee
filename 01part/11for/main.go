package main

import "fmt"

//for循环

func main() {
	////基本格式
	//for i := 0; i < 10; i++ {
	//	fmt.Println(i)
	//}
	//
	////变种1
	//var i = 5
	//for ; i < 10; i++ {
	//	fmt.Println(i)
	//}
	////变种2
	//var i0 = 5
	//for i0 < 10 {
	//	fmt.Println(i0)
	//	i0++
	//}
	//
	////无线循环
	////for{
	////fmt.Printf()
	//
	////for range循环
	//s := "hello"
	//for i, v := range s {
	//	fmt.Printf("%d %c\n", i, v)
	//}
	////哑元变量，不想用的变量用其表示
	//for _, v := range s {
	//	fmt.Printf("%c\n", v)
	//}
	//打印九九乘法表
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d\t", i, j, i*j)
		}
		fmt.Println()
		//fmt.Print(i)
		//fmt.Print(i,"\t")
		//fmt.Printf("%d\t",i)
	}

}
