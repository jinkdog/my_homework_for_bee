package main

import "fmt"

func main() {
	var mapSlice = make([]map[string]string, 3)
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
	fmt.Println("after init")
	// 对切片中的map元素进行初始化
	mapSlice[0] = make(map[string]string, 10)
	mapSlice[0]["name"] = "小王子"
	mapSlice[0]["password"] = "123456"
	mapSlice[0]["address"] = "沙河"
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}

	var ms = make([]map[string]string, 10) //前后map的类型要一致
	fmt.Println(ms)
	for index, value := range ms {
		//fmt.Printf("index:%dvalue:%d", index, value)
		fmt.Printf("index:%dvalue:%v\n", index, value)
	}
	//ms[0]=make([string]string,10)
	ms[0] = make(map[string]string, 10)
	ms[9] = make(map[string]string, 10)
	//ms[0]=["abc"]"abc"
	//ms[0]=["qwqe"]"qwqe"
	//ms[1]=["xyz"]"xyz"
	//ms[1]=["asad"]"asad"格式错误
	ms[0]["zxc"] = "wsd"
	ms[0]["qwe"] = "asd"
	ms[9]["num"] = "123"
	ms[9]["456"] = "456"
	//for index,value=range ms{
	//	fmt.Printf("index%dvalue%d\n",index,value)
	//}
	for index, value := range ms {
		fmt.Printf("index%dvalue%v\n", index, value)
	}
	//先初始化slice后初始化map
}
