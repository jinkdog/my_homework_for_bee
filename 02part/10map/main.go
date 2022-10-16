package main

//声明map，初始化开辟内存空间，并建立几个map类型的数据
//判断某个key是否存在
import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {

	//var m1=map[string]int没有等号
	var m1 map[string]int
	m1 = make(map[string]int, 10)
	//m1=map["a"]10错误使用
	m1["zfy"] = 10
	m1["yfz"] = 11
	fmt.Println(m1["zfy"])
	fmt.Printf("%T", m1["zfy"])
	m2 := map[string]int{
		"a": 12, "b": 13,
	}
	fmt.Println(m2["a"])
	fmt.Println(m1["zsd"]) //如果不存在拿到对应key的0值
	v, k := m1["zxc"]
	//if !=true{
	//	fmt.Println(k),
	//}else
	//{
	//	fmt.Println(v)
	//}if的条件书写错误，应判断k

	if k {
		fmt.Println("no")
	} else {
		fmt.Println(v)
	}
	v1, k1 := m1["zxc"]
	if !k1 {
		fmt.Printf("yes")
	} else {
		fmt.Println(v1)
	}

	//map的遍历

	//v2,k2:=m1 in range
	//for v2,k2:=m1 in range{
	//	fmt.Println(m1)
	//}不对
	for v2, k2 := range m1 {
		fmt.Println(v2, k2)
	}
	for k3 := range m1 {
		fmt.Println(k3)
	}
	for k2 := range m1 {
		fmt.Println(k2)
	}
	for _, v2 := range m1 {
		fmt.Println(v2)
	}
	//delete(m1["yfz"]11)
	//格式错误
	delete(m1, "yfz")
	fmt.Println(m1)
	//rand.seed(time().UnixNano())
	//rand.Seed(time.Now()UnixNano())
	rand.Seed(time.Now().UnixNano())
	m2 = make(map[string]int, 100)
	//for i=0;i<10;i++{
	//	key:=fmt.Sprintf("abc",10)
	//	value:=
	//}
	//for i := 0; i < 10; i++ {
	//	key := fmt.Sprintf("abc%d", 10)
	//	value := rand.Intn(10)
	//
	//}
	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("abc%02d", i)
		value := rand.Intn(100)
		m2[key] = value
	}
	//var keys=make([string],10)
	//for key range{
	//	key=keys
	//}
	//var keys=make([]string,0,10)
	//for key range{
	//	keys=append(keys,key)
	//}
	//var keys = make([]string, 10)缺少给string类型的位置
	//for key := range m2 {
	//	keys = append(keys, key)
	//}
	//Sort(keys)
	//for _,key range m2{为什么这么写？
	//	fmt.Println(key,m2[key]),
	//}
	//var keys =make([]string,0,10)
	//for key:=range keys{
	//	keys=append(keys,key)
	//}
	var keys = make([]string, 0, 10)
	for key := range m2 {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	for _, key := range keys {
		fmt.Println(key, m2[key])
	}
	//rand.Seed(time.Now().UnixNano()) //初始化随机数种子
	//
	//var scoreMap = make(map[string]int, 200)
	//
	//for i := 0; i < 100; i++ {
	//	key := fmt.Sprintf("stu%02d", i) //生成stu开头的字符串,最后生成的为stu01.。。stu99
	//	value := rand.Intn(100)          //生成0~99的随机整数
	//	scoreMap[key] = value
	//}
	////取出map中的所有key存入切片keys
	//var keys = make([]string, 0, 200)
	//for key := range scoreMap {
	//	keys = append(keys, key)
	//}
	////对切片进行排序
	//sort.Strings(keys)
	////按照排序后的key遍历map
	//for _, key := range keys {
	//	fmt.Println(key, scoreMap[key])
	//}

}
