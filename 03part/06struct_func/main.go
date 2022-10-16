package main

//构造函数 用new开头
import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string, age int) person {
	//name:name
	//age:age
	//返回值为person结构体
	return person{
		name: name,
		age:  age,
	}
}

// 结构体是值类型
// 构造函数是结构体还是结构体指针
// 结构体庞大时拷贝指针，不大时拷贝值
func main() {
	a := newPerson("q", 11)
	b := newPerson("w", 12)
	fmt.Println(a)
	fmt.Println(b)
}
