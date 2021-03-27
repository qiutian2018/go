package main
import (
	"fmt"
	"time"
)

func main(){
	var(
		num1,num2,num3 = 1,2,3
	)
	fmt.Println(num1,num2,num3)

	var num uint64 = 65535
	size := 64
	fmt.Printf("类型为uint64的证书%d需占用的存储空间为%d个字节\n",num,size)

	var num4 = -0x1000
	fmt.Printf("十六进制%X表示的是%d\n",num4,num4)

	var float = 5.89E-4
	fmt.Printf("浮点数%E表示的是%f\n",float,float)

	var numbers2 [5]int
		numbers2[0] = 2
		numbers2[3] = numbers2[0] - 3
		numbers2[1] = numbers2[2] + 5
		numbers2[4] = len(numbers2)
		sum := 11
		fmt.Printf("%v\n",(sum == numbers2[0]+numbers2[1]+numbers2[2]+numbers2[3]+numbers2[4]))
	
	var numbers3 = [7]int{1,2,3,4,5,6,7}
		slice3 := numbers3[1:len(numbers3)]
		length := 4
		capacity := 5
		fmt.Printf("%v,%v\n,%d,%d,%d\n",(length == len(slice3)),(capacity == cap(slice3)),len(numbers3),len(slice3),cap(slice3))
		fmt.Printf("%d,%d,%d,%d\n",slice3[0],slice3[1],slice3[2],slice3[3])

	ch2 := make(chan string, 1)
	go func(){
		ch2 <- "已到达!"
		ch2 <- "已到达!"
	}()
	var value string = "数据"
	s,ok := <- ch2
	s2,ok2 := <- ch2
	// s,ok := <- ch2
	// s,ok := <- ch2
	// s,ok := <- ch2
	// s,ok := <- ch2
	// s,ok := <- ch2
	value = value + (s)
	fmt.Println(value,s,ok,s2,ok2)

	type Sender chan<- int
	type Receiver <-chan int
	var myChannel = make(chan int, 0)
	var number = 6
	go func(){
		var sender Sender = myChannel
		sender <- number
		fmt.Println("Sent!")
	}()
	go func(){
		var receiver Receiver = myChannel
		fmt.Println("Receiverd!",<-receiver)
	}()
	time.Sleep(time.Second)
}