package main

import "fmt"

//运算符

func main() {
	var (
		a = 5
		b = 2
	)
	//算术运算符
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)
	a++ //表示一个单独语句，不能放在等号右边==>a=a+1
	b--
	fmt.Println(a, b)

	//关系运算符
	fmt.Println(a == b) //Go语言是强类型，相同类型的变量才能比较
	fmt.Println(a != b) //不等于
	fmt.Println(a >= b) //大于等于
	fmt.Println(a > b)  //大于
	fmt.Println(a <= b) //小于等于
	fmt.Println(a < b)  //小于

	//逻辑运算符
	//如果年龄大于18岁且年龄小于60岁
	age := 22
	if age > 18 && age < 60 {
		fmt.Println("work")
	} else {
		fmt.Println("不用上班")
	}
	//如果年龄小于18或大于60
	if age < 18 || age > 60 {
		fmt.Println("work")
	} else {
		fmt.Println("不用上班")
	}
	//not取反，原来为真就为假，原来为假就为真
	ismarried := false
	fmt.Println(ismarried)  //false
	fmt.Println(!ismarried) //true

	//位运算：针对的是二进制数
	//5的二进制表示：101
	//2的二进制表示：10

	//&：按位与（两位均为1才为1）
	fmt.Println(5 & 2) //000
	//|：按位或（两位有1个为1就为1）
	fmt.Println(5 | 2) //111
	//^:按位异或（两位不一样则为1）
	fmt.Println(5 ^ 2) //111
	//<<：将二进制位左移指定位置
	fmt.Println(5 << 1) //1010
	//>>：将二进制位右移指定位置
	fmt.Println(5 >> 1)
	//移动范围不要超出变量范围

	//赋值运算符，用来给变量赋值

	var x int
	x = 10
	x += 1 //x=x+1
	x -= 1 //x=x-1
	x *= 1 //x=x*2
	x /= 1 //x=x/2
	x %= 1 //x=x%

	x <<= 2 //x=x<<2
	x &= 2  //x=x&2
	x |= 3  //x=x|3
	x ^= 4  //x=x^4
	x >>= 2 //x=x>>2

	fmt.Println(x)
}
