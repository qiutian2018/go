package main
import (
	"fmt"
)

func printSlice(s []int){
	fmt.Printf("len=%d,cap=%d,val=%d,type=%T\n",len(s),cap(s),s,s)
}

func Pic(dx,dy int) [][]uint8{
	slice := make([][]uint8,dy)
	fmt.Println(slice,len(slice))
	for x := range slice{
		b := make([]uint8,dx)
		for y := range b{
			b[y] = uint8(x*y)
		}
		slice[x] = b
		// fmt.Println(b,len(b))
		// fmt.Println(slice,len(slice))
	}
	return slice
}

func main(){
	var a1 [3]int
	a1[1] = 1
	a1[2] = 2
	fmt.Println(a1)

	a2 := [7]int{0,2,2,3,4,5,6}
	fmt.Println(a2)
	
	s := a2[1:5]
	var s2 []int = a2[1:4]
	s2[2] = 333
	fmt.Println(s,s2,a2)

	s3 :=[]struct {
		name string
		addr string
	}{
		{"张三","北京"},
		{"李四","上海"},
		{"王五","aaa"},
	}

	fmt.Println(s3,s3[0].name)

	aa :=[]int{1,2,3,4,5}
	printSlice(aa)
	printSlice(aa[1:len(aa)-2])
	fmt.Println("---")
	printSlice(aa[:5])
	printSlice(aa[2:])
	printSlice(aa[:])
	printSlice(aa[0:])
	printSlice(aa[:0])
	n := aa[:2]
	printSlice(n)
	n = aa[:4]
	printSlice(n)
	n = aa[2:]
	printSlice(n)
	var t []int
	t2 := t
	printSlice(t)
	if t==nil{
		fmt.Println("t is nil")
	}
	if t2==nil{
		fmt.Println("t2 is nil")
	}

	var t3 []int = []int{1,2,3}
	t4 := t3[:0]
	t5 := t3[3:]
	t6 := t3
	printSlice(t3)
	printSlice(t4)
	printSlice(t5)
	if t4==nil{
		fmt.Println("t4 is nil")
	}
	if t5==nil{
		fmt.Println("t5 is nil")
	}
	if t6==nil{
		fmt.Println("t6 is nil")
	}
	// fmt.Printf("%b\n",t3==t3)
	fmt.Printf("%v\n",a2==a2)

	var tt [3]int = [3]int{2,3,4}
	var tt2 []int=[]int{2,3,4}
	var tt3 [3]int
	tt21 := make([]int,3)
	fmt.Printf("%T\n",tt)
	// if tt==nil{
	// 	fmt.Println("tt is nil")
	// }
	if tt2==nil{
		fmt.Println("tt2 is nil")
	}else{
		fmt.Println("tt2 is not nil")
	}
	if tt==[3]int{2,3,4} {
		fmt.Println("tt = [3]int{2,3,4}")
	}
	fmt.Printf("%T\n",tt3)
	if tt3 == [3]int{0,0,0}{
		fmt.Println("tt3 == []int{0,0,0}")
	}
	printSlice(tt21)
	if tt21 == nil {
		fmt.Println("tt21 is nil")
	}else{
		fmt.Println("tt21 is not nil")
		// fmt.Printf("%v\n",tt21==[]int{0,0,0})
	}

	fmt.Println("****1***")
	// ak := make([]int,5)
	ak := make([]int,0,5)
	printSlice(ak)
	fmt.Println(len(ak),cap(ak))
	ak = ak[:cap(ak)]
	printSlice(ak)
	ak = ak[1:]
	printSlice(ak)

	fmt.Println("*********2*********")
	var slice []int
	printSlice(slice)
	slice = append(slice,02,3,3)
	printSlice(slice)


	slice2 := []int{1,2,3,4,5,6,7,8}
	printSlice(slice2)

	for i,v:=range slice2{
		fmt.Printf("index=%d,val=%d\n",i,v)
	}

	fmt.Println(Pic(2,3))
	// pic.Show(Pic(2,3))
	// printSlice(Pic(2,8))
}