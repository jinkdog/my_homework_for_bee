package main

import "fmt"

func main() {
	//跳出多层for循环
	//var flag =false
	//for i := 0; i < 10; i++ {
	//	for j := 'A'; j < 'Z'; j++ {
	//		if j == 'D' {
	//			flag = true
	//			break //跳出内层循环
	//		}
	//		fmt.Printf("%v-%v\n", i, j)
	//	}
	//	if flag {
	//		break //跳出for循环（外层for循环）
	//	}
	//}
	for i := 0; i < 10; i++ {
		for j := 'A'; j < 'Z'; j++ {
			if j == 'D' {
				goto XX //跳到指定位置标签
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
XX: //labe标签
	fmt.Println("over")
}
