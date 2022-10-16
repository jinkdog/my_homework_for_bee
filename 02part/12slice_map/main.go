package main

import "fmt"

func main() {
	//var sliceMap = make(map[string][]string, 3)
	//fmt.Println(sliceMap)
	//fmt.Println("after init")
	//key := "中国"
	//value, ok := sliceMap[key]
	//if ok {
	//	value = make([]string, 0, 2)
	//}
	//value = append(value, "北京", "上海", "g")
	//sliceMap[key] = value
	//fmt.Println(sliceMap)

	//var ms=map[string][]string,3忘记make的初始化使用
	var sm = make(map[string][]string, 11)
	fmt.Println(sm)
	key := "abc"
	//value,ok=sm
	//value,ok=sm[key]
	//if !=value{
	//
	//}
	value := make([]string, 0, 3)
	value = append(value, "b", "c")
	sm[key] = value
	fmt.Println(sm)
	//先初始化map后初始化slice
}
