package main

import "fmt"

// if条件判断
func main() {
	//age := 19
	//if age > 18 {
	//	fmt.Println("sex")
	//} else {
	//	fmt.Println("no sex")
	//}
	//
	////多条件判断
	//if age>35{
	//	fmt.Println(("old")
	//}else if age>18{
	//	fmt.Println("young")
	//}else{
	//	fmt.Println("study")
	//}
	//作用域
	//age变量此时只在if条件判断语句中生效
	if age := 19; age > 18 {
		fmt.Println("sex")
	} else {
		fmt.Println("no sex")
	}
	//fmt,Println(age)无法打印
}
