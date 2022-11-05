package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strings"
	"sync"
)

var 神之心 sync.Map

func alignSize(num []int) int {
	size := 0
	for _, n := range num {
		if s := int(math.Log10(float64(n))); s > size { //取num数组中输入的value来取10对数来确定数字位数，再将得出的位数赋值给size
			size = s
		}
	}
	return size
}

func main() {
	//map的使用

	//初始化

	//方法1：可以先用make初始化后，再给变量赋值
	//赛博朋克 := make(map[string]string)
	//赛博朋克 = map[string]string{
	//	"大卫":  "传奇",
	//	"露西":  "寡妇",
	//	"瑞贝卡": "瑞贝卡酱",
	//}
	//fmt.Println(赛博朋克)
	//方法2：可以直接声明map变量的值，不用make初始化
	//赛博朋克 := map[string]string{
	//	"大卫":  "传奇",
	//	"露西":  "寡妇",
	//	"瑞贝卡": "瑞贝卡酱",
	//}
	//方法3：逐个赋值
	//赛博朋克 := make(map[string]string)
	//赛博朋克["大卫"] = "传奇"
	//赛博朋克["露西"] = "寡妇"
	//赛博朋克["瑞贝卡"] = "瑞贝卡酱"
	//fmt.Println(赛博朋克)
	//实操结论:map关键词不能被替代（想办法把它替代或者在打印时不显示它）

	//空map实现
	//a:=map[string]int{}
	//fmt.Println(a)

	//访问
	//a1 := map[string]int{
	//	"a": 1,
	//}
	//fmt.Println(a1["a"])

	//go语言中读取一个不存在的key
	//type Node struct {
	//	Next  *Node
	//	Value interface{}
	//}
	//
	//var first *Node
	//
	//func main(){
	//	visited := make(map[*Node]bool)
	//	for n := first; n != nil; n = n.Next {
	//		if visited[n] {
	//			fmt.Println("cycle detected")
	//			break
	//		}
	//		visited[n] = true
	//		fmt.Println(n.Value)
	//	}
	//}

	//删除
	//func delete(m map[Type]Type1, key Type)delete函数源码
	//delete(赛博朋克, "荒板")
	//fmt.Println(赛博朋克)
	//delete(赛博朋克, "露西")
	//fmt.Println(赛博朋克)
	//实操结论：可以删除不存在的键

	//赋值
	//赛博朋克["大卫"] = 赛博朋克["赛博疯子"]
	//fmt.Println(赛博朋克)
	////有趣的发现，如果把map中存在的key值换成map中不存在的key值，打印出的key不会改变，但是value会变为nil
	//赛博朋克["大卫"] = "传奇"
	//fmt.Println(赛博朋克)                      //赛博朋克["大卫"]说明这个已经代表的就是一个value
	//赛博朋克["大卫"] = 赛博朋克["大卫"] + " 生的伟大，死的光荣" //可以直接在字符串后加字符串
	//fmt.Println(赛博朋克)
	//

	////遍历
	//for name, character := range 赛博朋克 {
	//	fmt.Println(name, character)
	//}
	////fmt.Println(name,character)只能在for循环内接受name和character
	//for name, _ := range 赛博朋克 {
	//	fmt.Println(name)
	//}
	//用匿名变量接收

	////线程安全
	////1slice
	//var s0 []string
	//for i := 0; i < 50; i++ {
	//	go func() {
	//		s0 = append(s0, "我明白")
	//	}()
	//}
	//
	//fmt.Println(s0)
	//fmt.Printf("我明白了%d次\n", len(s0))
	////疑问：为什么我明白的个数和执行的次数不一样？
	//
	//var s1 []int
	//for i := 0; i < 50; i++ {
	//	go func() {
	//		s1 = append(s1, i)
	//	}()
	//}
	//fmt.Println(s1)
	//fmt.Printf("运行了%d次\n", len(s1))
	////s1中的数字个数也和运行次数不同
	//
	////2map
	//s2 := make(map[string]string)
	//for i := 0; i < 20; i++ {
	//	go func() {
	//		s2["电次"] = "那由多"
	//	}()
	//}
	//fmt.Printf("拥抱了%d次\n", len(s2))
	////fatal error: concurrent map writes

	//如何保证并发安全

	//写入 Store写入一个key值

	//神之眼 := []string{"风", "岩", "雷", "草", "水", "火", "冰"}
	//for i := 0; i < 6; i++ {
	//	go func() {
	//		神之心.Store(i, 神之眼[i])//由 'go' 语句中的 'func' 文字捕获的循环变量可能有意外值//这个为什么不能算闭包
	//	}()
	//}
	//time.Sleep(time.Second)

	//神之眼 := []string{"风", "岩", "雷", "草", "水", "火", "冰"}
	//for i := 0; i < 7; i++ {
	//	go func(i int) {
	//		神之心.Store(i, 神之眼[i])
	//	}(i)
	//}
	//time.Sleep(time.Second)
	////读取 Load返回存储在 map 中的键的值，如果没有值，则返回 nil。ok 结果表示是否在 map 中找到了值。
	//v, ok := 神之心.Load(6)
	//fmt.Printf("Load:%v,%v\n", v, ok)
	////删除 删除某一个键的值。
	//神之心.Delete(1)
	////读或写
	//v, ok = 神之心.LoadOrStore(1, "域") //如果key不存在，储存值返回false
	////v, ok = 神之心.LoadOrStore(0, "金") //如果key存在，那么value不会改变,返回true
	//fmt.Printf("LoadOrStore:%v,%v\n", v, ok)
	//fmt.Println(神之心.Load(1))
	//fmt.Println(神之心)
	////只打印神之眼的话，对应的值不会变，实际上是把神之眼的值拷贝到神之心中，并赋予key值
	//
	////读或删  LoadAndDelete：删除一个键的值，如果有的话返回之前的值。
	//v, ok = 神之心.LoadAndDelete(1)
	//fmt.Printf("LoadAndDelete:%v,%v\n", v, ok)
	//
	////  Range：递归调用，对 map 中存在的每个键和值依次调用闭包函数 f。如果 f 返回 false 就停止迭代。
	//神之心.Range(func(key, value interface{}) bool {
	//	fmt.Printf("%v,%v\n", key, value)
	//	return true
	//})
	////疑问：为什么打印出来的顺序不一样？
	//

	////fmt包

	////print系列 略
	////sprint系列，用来让变量提前打印好的数据
	//year, month, day := time.Now().Date()
	//hour, minute, second := time.Now().Clock()
	//s3 := fmt.Sprintf("此刻%d,%d,%d,%d,%d,%d", year, month, day, hour, minute, second) //接收数字
	//s4 := fmt.Sprintln(year, month, day, hour, minute, second)                       //接收时刻
	//fmt.Println(s3, s4)
	//fprint系列，可以用来直接编写指定文件的数据
	//var a string = "test Fprint"
	//fmt.Fprintln(os.Stdout, a)
	////Errorf函数
	//error := "哼唧"
	//err := fmt.Errorf("出错的话就%s一声", error)
	//fmt.Println(err)
	//fmt.Println(error)
	//fmt.Printf("出错的话就%s一声", error)
	//err是error类型的数据
	//疑问1：error类型的数据有什么用，怎么用？
	//疑问2:不明白errorf这样赋值有什么意义

	//输入
	//a, b, c := true, 10, false
	//fmt.Println(a, b, c)
	//fmt.Scan(&a, &b, &c)
	//1可以全部输入后再回车，中间用空格衔接
	//2可以先输入a，b再回车，输入c再回车=》可以没有固定顺序地输入
	//3字符串型输入true返回true单词，整型输入true停止输入返回原值，布尔型输入true继续输入
	//fmt.Println(a, b, c)

	//Scanln
	//d1, b1, c1 := "", 11, false
	//fmt.Println(d1, b1, c1)
	//fmt.Scanln(&d1, &b1, &c1)
	////1Scanln是一次性输入完再按回车，否则只能输入其中一个元素
	////2字符串型输入true返回true单词，整型输入true停止输入返回原值，布尔型输入true继续输入
	//fmt.Println(d1, b1, c1)

	//格式化字符串指定宽度
	//d2, b2, c2 := " ", 12, false
	//fmt.Println(d2, b2, c2)
	//fmt.Scanf("%s%1d%4t", &d2, &b2, &c2)
	//Scanf也是只能按一次回车算输入
	//%s和%d可以在格式化符之前加入数字，%t若对应的是布尔型则只能为数字4及以上输入方可能改变数字，否则无法改变
	//fmt.Println(d2, b2, c2)

	//格式化符
	//通用
	//type 大学 struct {
	//	双非 string
	//}
	//油砖 := 大学{
	//	"重庆邮电大学",
	//}
	//
	//fmt.Printf("%+v\n", 油砖)
	//fmt.Printf("%#v\n", 油砖)
	//fmt.Printf("%T\n", 油砖)
	//fmt.Printf("%v的保研率%%0.00001\n", 油砖)
	//fmt.Printf("%+v\n%#v\n%T\n%v的保研率%%0.00001\n", 油砖, 油砖, 油砖, 油砖)
	//不能只输入一个油砖，要输入对应数量的油砖，否则，一个都输出不出来

	//整数
	//fmt.Printf("%b\n", 11)
	//fmt.Printf("%c\n", 0x4E3d)
	//fmt.Printf("%d\n", 0x12)
	//fmt.Printf("%o\n", 10)
	//fmt.Printf("%q\n", 0x4E3d)//多了个单引号
	//fmt.Printf("%x\n", 11)
	//fmt.Printf("%X\n", 11)
	//fmt.Printf("%U\n", 0x4E3d)

	//浮点数和复数的两个部分：
	//fmt.Printf("%b\n", 1.2) //完全不明白有什么用
	//fmt.Printf("%e\n", 1.222)
	//fmt.Printf("%f\n", 1.222000000000001) //会忽略浮点数位上限后的数字
	//fmt.Printf("%g\n", 1.2220000000001)
	//fmt.Printf("%G\n", 1.220003)

	//字符串和[]byte
	//fmt.Printf("%s\n", "aaa")
	//fmt.Printf("%q\n", "aaa")
	//fmt.Printf("%x\n", "aaa")
	//fmt.Printf("%X\n", "sopdcjps"+
	//	"") //按回车就可以实现

	//指针
	//fmt.Printf("%p", &油砖) //是小写的p，讲义错误，但不完全错
	//其他
	//不想搞感觉没什么应用场景

	//宽度
	//func Log10(x float64) float64 // 10 为底的对数函数
	//nums := []int{1, 11, 111, 1111}
	//size := alignSize(nums)
	//for i, n := range nums {
	//	fmt.Printf("%04d %*d \n", i, size, n)
	//}
	//fmt.Printf("%04(最后一位数根据i变化而变化，0用来补4位中剩余的3位)d %*（*指根据size来改变输入的宽度，即打印多少位n）d \n", i, size, n)
	//通过位置引用
	//fmt.Printf("你%[1]s我，我%[1]s你\n", "爱")

	//time包

	//时间类型
	//now := time.Now()
	//year := now.Year() //相当于time.Now().Year()
	//month := now.Month()
	//day := now.Day()
	//hour := now.Hour()
	//minute := now.Minute()
	//second := now.Second()
	//fmt.Println(year, month, day, hour, minute, second)

	//时间戳
	//获取当前时间的时间戳
	//不同单位累加到现在有多久
	//stt := now.Unix()
	//mstt := now.UnixMilli()
	//mmstt := now.UnixMicro()
	//nstt := now.UnixNano()
	//fmt.Println(stt, mstt, mmstt, nstt)

	//将时间戳转换为时间格式
	//var timestamp int64
	//timeobj := time.Unix(timestamp, 0)
	//year := timeobj.Year()
	//month := timeobj.Month()
	//day := timeobj.Day()
	//hour := timeobj.Hour()
	//minute := timeobj.Minute()
	//second := timeobj.Second()
	//fmt.Println(year, month, day, hour, minute, second)
	//fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
	//疑问：怎么打印将此刻的时间戳转换为此刻的时间？

	//时间间隔

	//Location和timezone
	//UTC世界协调时间
	//secondsEasrOfUTC := int((8 * time.Hour).Seconds())
	//beijing := time.FixedZone("Beijing Time", secondsEasrOfUTC)
	//newyork, err := time.LoadLocation("America/New_York")
	//if err != nil {
	//	fmt.Println("load America/New_York location failed", err)
	//	return
	//}
	//fmt.Println()
	//shanghai, err := time.LoadLocation("Asia/Shanghai")
	//tokyo, err := time.LoadLocation("Asia/Tokyo")
	//
	//timeInLocal := time.Date(2009, 1, 1, 12, 0, 0, 0, time.Local)
	//timeInUTC := time.Date(2009, 1, 1, 12, 0, 0, 0, time.UTC)
	//sameTimeInBeijing := time.Date(2009, 1, 1, 20, 0, 0, 0, beijing)
	//sameTimeInNewYork := time.Date(2009, 1, 1, 7, 0, 0, 0, newyork)
	//timesAreEqual := timeInUTC.Equal(sameTimeInBeijing)
	//fmt.Println(timesAreEqual)
	//timesAreEqual = timeInUTC.Equal(sameTimeInNewYork)
	//fmt.Println(timesAreEqual)
	////疑问：这仅仅是明白了时区的概念，实际上还是没办法运用，该如何使用？

	//时间操作
	//Add
	//now := time.Now()
	//later := now.Add(10000 * time.Hour)
	//fmt.Println(later)

	//Sub
	//now := time.Now()
	//later := now.Add(10000 * time.Hour)
	//tu := later.Sub(now)
	//fmt.Println(tu)
	//Sub只能是两个时间点相减得出一个时间段
	//now := time.Now()
	//tu := now.Add(-100000 * time.Second)
	//fmt.Println(tu)
	//可以用Add一个负时间段得出一个时间点

	//Equal
	//t := sameTimeInBeijing.Equal(sameTimeInNewYork)
	//fmt.Println(t)
	//shanghai, err := time.LoadLocation("Asia/Shanghai")
	//if err != nil {
	//	fmt.Println(err)
	//	return//return可有可无
	//}
	//sametimeshanghai := time.Date(3000, 0, 0, 0, 0, 0, 0, shanghai)
	//tokyo, err := time.LoadLocation("Asia/Tokyo")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//sametimetokyo := time.Date(3000, 0, 0, 0, 0, 0, 0, tokyo)
	//t := sametimetokyo.Equal(sametimeshanghai)
	//fmt.Println(t)

	//Before/After
	//shanghai, err := time.LoadLocation("Asia/Shanghai")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//timeshanghai := time.Date(3000, 0, 0, 0, 0, 0, 0, shanghai)
	//tokyo, err := time.LoadLocation("Asia/Tokyo")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//timetokyo := time.Date(3000, 0, 0, 0, 0, 0, 0, tokyo)
	//t1 := timeshanghai.After(timetokyo)
	//t2 := timeshanghai.Before(timetokyo)
	//fmt.Println(t1, t2)

	//Timer与Ticker
	//timer := time.NewTimer(10 * time.Second)
	//ticker := time.NewTimer(1 * time.Second)
	//defer ticker.Stop()
	//for {
	//	select {
	//	case <-timer.C:
	//		fmt.Println("Done")
	//		return
	//	case t := <-ticker.C:
	//		fmt.Println("ticker at:", t)
	//		//此处如果加入return则ticker当即返回程序结束
	//	}
	//}
	//疑问：如何实现在10秒内，ticker每隔1s冒出的效果？
	//具体使用还是没搞懂

	//时间格式化
	//now := time.Now()
	////24小时制
	//fmt.Println(now.Format("2006-1-2-3-4-5 Mon Jan "))
	////12小时制
	//fmt.Println(now.Format("2006-1-2-3-4-5 PM Mon Jan "))
	////格式化输出保留到小数点后三位，因为有0占位
	//fmt.Println(now.Format("2006/01/02 15:04:05.000"))
	////还是会输出保留小数点后三位，但是最后如果是0的话会被省略
	//fmt.Println(now.Format("2006/01/02 15:04:05.999"))
	////只格式化时刻
	//fmt.Println(now.Format("15:04:05.999"))
	////只格式化日期
	//fmt.Println(now.Format("2006/01/02"))

	//解析字符串格式的时间
	////time.Parse()解析时不用加入时区
	//timeobj, err := time.Parse("2006/01/02", "2022/11/02")
	////"2006/01/2", "2022/11/02"可以
	////"2006/01/02", "2022/11/2"不行
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(timeobj)
	//
	//timeObj, err := time.Parse(time.RFC3339, "2022-10-05T11:25:20+08:00")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(timeObj)

	//time.ParseInLocation（）需加入时区信息
	//now := time.Now()
	//fmt.Println(now)
	//
	//loc, err := time.LoadLocation("Asia/Shanghai")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//timeobj, err := time.ParseInLocation("2006/01/02", "2022/11/02", loc)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(timeobj)
	//fmt.Println(timeobj.Sub(now))

	//string包

	//包含判断

	//前后缀判断
	//前缀prefix 后缀suffix
	//var str string = "this is a cat"
	//var str2 string = "Hello Boss"
	//var str3 string = "Lucy is my name"
	//fmt.Printf("T/F? Dose the string \"%s\" have prefix %s?\n", str, "this ")
	//fmt.Printf("%t\n", strings.HasPrefix(str, "this "))
	////无论prefix向后延续多长都可以执行，只要中间不间断
	//fmt.Printf("T/F? Dose the string \"%s\" have prefix %s?\n", str2, "He")
	//fmt.Printf("%t\n", strings.HasPrefix(str2, "He"))
	//fmt.Printf("T/F? Dose the string \"%s\" have suffix %s?\n", str3, "my name")
	//fmt.Printf("%t\n", strings.HasSuffix(str3, "my name"))
	////无论suffix向前延续多长都可以执行，只要中间不间断
	//var 能做到吗 string = "德克萨斯能做到吗？"
	//fmt.Printf("%t\n", strings.HasPrefix(能做到吗, "德克萨斯"))

	//子字符串包含
	//ContainsAny查找字符
	//fmt.Println(strings.ContainsAny("team", "e"))
	//fmt.Println(strings.ContainsAny("chainsawman", " "))
	//fmt.Println(strings.ContainsAny("name", "n&m"))

	//Contains查找字符串
	//fmt.Println(strings.Contains("name", "n&m"))
	//fmt.Println(strings.Contains("name", "na"))
	//fmt.Println(strings.Contains("chain", "")) //所有字符串类型的都包含字符串空

	//索引
	//str := "I am groot."
	//fmt.Printf("The position of\"groot\"is:")
	//fmt.Printf("%d\n", strings.Index(str, "groot"))
	////如果打印不存在的字符串，Index返回-1
	//fmt.Printf("The position of\"Groot\"is:")
	//fmt.Printf("%d\n", strings.Index(str, "Groot"))
	//fmt.Printf("The position of the last of instance of\"Hi\"is:")
	//fmt.Printf("%d\n", strings.LastIndex(str, "I"))
	//fmt.Printf("%d\n", strings.LastIndex(str, "g"))
	//fmt.Printf("%d\n", strings.LastIndex(str, "."))
	//fmt.Printf("%d\n", strings.LastIndex(str, " "))
	////index的索引以一个字母为依据数“”是最后的最后一个
	//str2 := "我一定会回来的"
	//fmt.Printf("the positon of\"我\"is:")
	//fmt.Printf("%d\n", strings.IndexRune(str2, '我')) //非Ascii码字符得用单引号
	//fmt.Printf("%d\n", strings.IndexRune(str2, '回')) //一个汉字占3个位置

	//替换
	//str := "aser is the saber,saber is one of the 3 hero.saber is a character"
	//new := "lancer"
	//old := "saber"
	//n := 0 //n指替换几个,而不是替换第几个
	////n := -1//替换所有
	//println(strings.Replace(str, old, new, n))
	//fmt.Println(strings.Replace(str, old, new, n)) //和上面的没有差别

	//统计

	//出现频率
	//str := "Baby don't you know don't you love is so sweet"
	//str1 := "你问我爱你有多深，我爱你有几分。"
	//var how_many_o string = "o"
	//fmt.Printf("%d\n", strings.Count(str, how_many_o))
	//fmt.Printf("%d\n", strings.Count(str1, "你"))
	//fmt.Printf("%d\n", strings.Count(str1, "熊子瑜"))

	//字符数量
	//str := "原来你也玩原神"
	//fmt.Printf("%d\n", len([]rune(str)))
	//fmt.Println(utf8.RuneCountInString(str))

	//大小写转换
	//str := "Trail's fire you always know\n"
	//fmt.Printf("%s", str)
	//fmt.Printf(strings.ToLower(str))
	//fmt.Printf(strings.ToUpper(str))

	//修剪
	//str := "我心如铁 小鲤鱼 坚不可摧"
	//str := "AAA 小鲤鱼 AAA"
	//fmt.Println(str)
	//fmt.Println(strings.Trim(str, "小鲤鱼"))
	//fmt.Println(strings.TrimLeft(str, "小鲤鱼"))
	//fmt.Println(strings.TrimRight(str, "小鲤鱼"))
	//fmt.Println(strings.TrimSpace(str))

	//str := "袁神 Hello Go 袁神"
	//str1 := "          请使用TrimSpace()修剪空白字符       "
	//fmt.Println(str)
	//fmt.Printf("%q\n", strings.Trim(str, "袁神")) //修剪前缀和后缀
	//fmt.Printf("%q\n", strings.Trim(str, "袁神 "))
	//
	//fmt.Printf("%q\n", strings.TrimLeft(str, "袁神 "))  //修剪左边前缀的字符
	//fmt.Printf("%q\n", strings.TrimRight(str, " 袁神")) //修剪右边后缀的字符
	//fmt.Printf("%q\n", strings.TrimSpace(str1))       //修剪前缀和后缀的空白字符
	//被修剪的字符放入Trim函数中

	//分割
	//str := "是身如焰，从渴爱生"
	//fmt.Println(str)
	//fmt.Printf("%q\n", strings.Split(str, "是"))
	//fmt.Printf("%q\n", strings.Split(str, ""))

	//os
	//目录

	//创建文件
	//name := "\\Users\\张丰毅\\go\\src\\main.go"
	//creatfile(name)

	//创建目录
	//创建的单个目录
	//name := "\\Users\\张丰毅\\go\\src\\main"
	//creatdir(name)

	//创建的多级目录
	//name := "\\Users\\张丰毅\\go\\src\\main1"
	//creatdirall(name)

	//删除目录

	//删除单目录
	//name := "\\Users\\张丰毅\\go\\src\\main"
	//Removedir(name)
	//只能删除单目录，如果目录里面还有目录或其他文件就删掉

	//删除路径下所有目录
	//name := "\\Users\\张丰毅\\go\\src\\main1"
	//Removeairall(name)

	//获取当前目录
	//GetWd()

	//修改当前目录
	//Chwd()

	//获取临时文件
	//Temdir()

	// 修改文件名
	//RenameFile("04", "03")//应该要加文件路径但是我懒得敲文件名字

	//文件读操作
	//打开文件
	//默认打开方式
	//Open("C:\\Users\\张丰毅\\go\\src\\github.com\\example\\go_test_for_gin\\main.go")
	//疑问：实际上并没有打开文件，不知道是怎么回事

	//以指定权限打开文件
	//OpenFile()
	//疑问：不知道该填入什么参数，不知道填入的参数从哪里获取

	//关闭文件
	//ReadFile01("C:\\Users\\张丰毅\\go\\src\\github.com\\studygo_for_LWZ\\06part\\02\\main.go")
	//ReadFile01("github.com/studygo_for_LWZ/06part/02/main.go")
	//疑问：文件并没有从上方消失，还是不知道如何使用

	//读取文件
	//ReadFile02("github.com/studygo_for_LWZ/06part/02/main.go")
	//疑问：不明白读取、打开的区别

	//读取目录
	//ReadDir("C:\\Users\\张丰毅\\go\\src\\awesomeProject\\homework_for_redrock")
	//疑问：实际运行后实现了”正在读取“然而之后无事发生

	//设置偏移量
	//Seek()
	//疑问：不知道偏移量是什么

	//获取文件信息
	//StateFile("C:\\Users\\张丰毅\\go\\src\\awesomeProject\\homework_for_redrock")
	//终于有成功实现的了

	//数据输入
	//WriteFile("C:\\Users\\张丰毅\\go\\src\\awesomeProject\\homework_for_bee\\04LV0\\main.go")
	//成功实现

	//字符串输入
	//WritrStringFile("C:\\Users\\张丰毅\\go\\src\\awesomeProject\\homework_for_bee\\04LV0\\main.go")

	//写入指定位置
	//WriteFileAt("C:\\Users\\张丰毅\\go\\src\\awesomeProject\\homework_for_bee\\04LV0\\main.go")

	//使用缓冲区
	//WriteFile02("C:\\Users\\张丰毅\\go\\src\\awesomeProject\\homework_for_bee\\04LV0\\main.go")

	//逐行读取
	//_ = ReadLine("C:\\Users\\张丰毅\\go\\\\src\\awesomeProject\\homework_for_bee\\04LV0\\main.go")
}
func creatfile(name string) {
	file, _ := os.Create(name)
	fmt.Println(file)
}

// 创建目录
func creatdir(name string) {
	_ = os.Mkdir(name, os.ModePerm)
}

// 创建多级目录
func creatdirall(name string) {
	_ = os.MkdirAll(name, os.ModePerm)
}

// 删除目录
func Removedir(name string) {
	_ = os.Remove(name)
}

// 删除路径下所有目录
func Removeairall(name string) {
	_ = os.RemoveAll(name)
}

// 获取当前目录
func GetWd() {
	dir, _ := os.Getwd()
	fmt.Println(dir)
}

// 改变当前工作目录
func Chwd() {
	_ = os.Chdir("")
}

// 获取临时文件
func Temdir() {
	fmt.Println(os.TempDir())
}

// 修改文件名
func RenameFile(oldpath, newpath string) {
	_ = os.Rename(oldpath, newpath)
}

// Open只读方式不能写
func Open(name string) {
	file, _ := os.Open(name)
	fmt.Println(file)
}

// 以指定权限打开文件
func OpenFile(name string, perm os.FileMode) {
	file, _ := os.OpenFile(name, os.O_RDONLY, perm)
	fmt.Println(file)
}

// 关闭文件
func ReadFile01(name string) {
	file, _ := os.Open(name)
	defer file.Close()

	buf := make([]byte, 100)
	n, _ := file.Read(buf)
	fmt.Println(n)
}

// 读取文件
func ReadFile02(name string) {
	file, _ := os.Open(name)
	fmt.Println(file)
	for {
		buf := make([]byte, 5) //设置缓冲区，一次读取5个字节
		//虽然我也不知道缓冲区是什么怎么用
		n, err := file.Read(buf)
		fmt.Println(string(buf))
		fmt.Println(n)
		if err == io.EOF { //表示文件读取完毕
			break
		}
	}
	file.Close()
}

// 从指定位置读取文件
func Filereadat(name string) {
	file, err := os.Open(name)
	if err != nil {
		fmt.Println(err)
	}

	buf := make([]byte, 10)
	n, err := file.ReadAt(buf, 5)
	fmt.Println(string(buf))
	fmt.Println(n)
	file.Close()
}

// 读取目录
func ReadDir(name string) {
	file, err := os.Open(name)
	if err != nil {
		fmt.Println(err)
	}
	dir, err := file.ReadDir(-1)
	if err != nil {
		fmt.Println(err)
	}
	for key, value := range dir {
		fmt.Println(key, value)
	}
}

// 设置偏移量
func Seek(name string) {
	file, err := os.Open(name)
	if err != nil {
		fmt.Println(err)
	}
	buf := make([]byte, 5)
	file.Seek(3, 0)
	n, err := file.Read(buf)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(buf))
	fmt.Println(n)
	file.Close()
}

// 获取文件信息
func StateFile(name string) {
	file, err := os.Open(name)
	if err != nil {
		fmt.Println(err)
	}
	finfo, err := file.Stat()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("f是否是个文件:%v\n", finfo.IsDir())
	fmt.Printf("f文件修改时间:%v\n", finfo.ModTime().String())
	fmt.Printf("f文件的名称:%v\n", finfo.Name())
	fmt.Printf("f文件的大小:%v\n", finfo.Size())
	fmt.Printf("f文件的权限:%v\n", finfo.Mode().String())
}

// 数据写入
func WriteFile01(name string) {
	file, err := os.OpenFile(name, os.O_RDWR, 0775)
	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i < 10; i++ {
		file.Write([]byte(fmt.Sprint("不要回答", i)))
	}
	file.Close()
}

// 字符串写入
func WritrStringFile(name string) {
	file, err := os.OpenFile(name, os.O_RDWR, 0775)
	if err != nil {
		fmt.Println(err)
	}
	file.WriteString("最好的东西")
}

// 写入指定位置
func WriteFileAt(name string) {
	file, err := os.OpenFile(name, os.O_RDWR, 0775)
	if err != nil {
		fmt.Println(err)
	}
	file.WriteAt([]byte("学？学个屁"), 10)
}

// 使用缓冲区
func WriteFile02(name string) {
	file, _ := os.OpenFile(name, os.O_RDWR, 0775)
	defer file.Close()
	writefile := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writefile.WriteString(fmt.Sprintf("这是第次%d循环\n", i))
	}
	writefile.Flush()
}

// 逐行读取
func ReadLine(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	buf := bufio.NewReader(f)

	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		fmt.Printf("数据行：%v", line)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	}

}
