package main

import (
	"fmt"
	"strings"
)

//go语言中字符串用双引号，字符用单引号
//字符串

func main() {
	//\本来具有特殊含义\\表示单纯的\
	path := "\"C:\\Users\\张丰毅\\Desktop\\C++\""
	path01 := "'C:\\Users\\张丰毅\\Desktop\\C++'"
	fmt.Println(path)
	fmt.Println(path01)

	s := "p"
	fmt.Println(s)
	//多行的字符串，esc下方键
	s2 := `
    我爱你
`
	fmt.Println(s2)
	s3 := `C:\Users\张丰毅\go\src\github.com\studygo`
	fmt.Println(s3)

	//字符串相关操作
	fmt.Println(len(s3))

	//字符串拼接
	name := "make"
	world := "love"

	ss := name + world
	fmt.Println(ss)
	fmt.Printf("%s%s", name, world)
	ss1 := fmt.Sprintf("%s%s", name, world)
	fmt.Println(ss1)
	//分割
	ret := strings.Split(s3, "\\")
	fmt.Println(ret)
	//包含
	fmt.Println(strings.Contains(ss, "zfy"))
	fmt.Println(strings.Contains(ss, "yfz"))
	//前缀
	fmt.Println(strings.HasPrefix(ss, "zfy"))
	//后缀
	fmt.Println(strings.HasSuffix(ss, "zfy"))

	s4 := "abcfbbe"
	fmt.Println(strings.Index(s4, "c"))
	fmt.Println(strings.LastIndex(s4, "b"))
	//拼接
	fmt.Println(strings.Join(ret, "+"))
}
