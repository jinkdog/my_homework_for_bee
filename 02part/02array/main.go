package main

import "fmt"

//数组

// 存放元素容器
// 必须指定存放元素的类型和长度
// 长度和容量是数组类型的一部分
func main() {
	var a1 [3]bool //[true f t]
	var a2 [4]bool //[ t t f  f]
	fmt.Printf("a1:%T a2:%T", a1, a2)
	//数组初始化
	//如果不初始化，默认数值为零值（布尔型：false 整型和浮点型：0 字符型：”“）
	//1 初始化方式1
	a1 = [3]bool{true, true, true}
	fmt.Println(a1)
	//2 初始化方式2
	a3 := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(a3)
	//3 初始化方式3 根据索引来初始化
	a4 := [6]int{0: 1, 5: 7}
	fmt.Println(a4)

	//数组的遍历
	city := [4]string{"北", "上", "广", "深"}
	//1 根据索引遍历
	for i := 0; i < len(city); i++ {
		fmt.Println(city[i])
	}
	//2 for range遍历
	for i, v := range city {
		fmt.Println(i, v)
	}

	//多维数组
	//[1,2],[3,4],[5,6]
	var a11 [3][2]int
	a11 = [3][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
		[2]int{5, 6},
	}
	fmt.Println(a11)

	//多维数组的遍历

	a := [3][2]string{
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	}
	for _, v1 := range a {
		for _, v2 := range v1 {
			fmt.Printf("%s\t", v2)
		}
		fmt.Println()
	}

	//数组是值类型,值传递后原来值不会变
	b1 := [3]int{1, 2, 3}
	b2 := b1
	b2[0] = 100
	fmt.Println(b1, b2)
}
