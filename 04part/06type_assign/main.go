package main

import "fmt"

// 类型接口1
func assign(a interface{}) {
	fmt.Printf("%T\n", a)
	str, ok := a.(string)
	if !ok {
		fmt.Printf("false")
	} else {
		fmt.Println(str)
	}
}

// 类型接口2
func assign2(a interface{}) {
	fmt.Printf("%T\n", a)
	switch t := a.(type) {
	case string:
		fmt.Println(t)
	case int:
		fmt.Println(t)
	case bool:
		fmt.Println(t)
	case int64:
		fmt.Println(t)
	}

}
func main() {
	assign(100)
	assign2(true)
}
