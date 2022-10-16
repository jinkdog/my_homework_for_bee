package main

import "fmt"

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	a := 1
	b := 2
	defer calc("1", a, calc("2", a, b)) //1
	a = 0
	b = 1
	defer calc("3", b, calc("4", a, b)) //2
	//运用defer时
	//将参数a，b带入先执行函数内函数而后将外函数入栈
	//继续执接下来的语句
	//最后先执行2后执行1外函数，其中变量均在defer前带入了

}
