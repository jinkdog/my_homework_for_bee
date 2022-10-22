package main

import "fmt"

//interface：空接口
//interface{}：空接口类型

// 空接口做函数参数
func show(a interface{}) { //interface{}还是要有变量名a
	fmt.Printf("type:%T value:%v\n", a, a)
}
func main() {
	//var m1 map[]string interface{}//map([]string interface)//忘记map语法
	var m1 map[string]interface{}
	m1 = make(map[string]interface{}, 16)
	m1["name"] = "zfy"
	m1["age"] = 18
	m1["merried"] = true
	m1["hobby"] = [...]string{"sing", "dance", "rap"} //数组间的元素用逗号隔开
	fmt.Println(m1)

	show(false)
	show(nil)
	show(m1)
}
