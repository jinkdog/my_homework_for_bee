package main

import "fmt"

func main() {
	ch1 := make(chan int, 100)
	ch1 <- 10
	ch1 <- 20

	close(ch1)
	//通道关闭后可以取值，但是只能取到对应类型的0值
	//取值方法1
	<-ch1
	<-ch1
	<-ch1
	//取值方法2
	for x := range ch1 {
		fmt.Println(x)
	}

}
