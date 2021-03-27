package main
import (
	"fmt"
	"math"
)

func compute(fn func(float64,float64)float64) float64{
	return fn(3,4)
}

func test(a,b int,s string) int{
	fmt.Println(s)
	return b
}

func addr() func (int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func fibonacci() func() int {
	bak1,bak2 := 0,1
	return func() int {
		tmp := bak1
		bak1,bak2 = bak2,(bak1+bak2)
		return tmp
	}
}

func fibonacci2(x int) int {
	if x==0 {
		return 0
	}
	if x==1 {
		return 1
	}
	return fibonacci2(x-1) + fibonacci2(x-2)
}

func main(){
	hypot := func(x,y float64) float64{
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(compute(hypot))

	fmt.Println(test(1,2,"sss"))

	pos,neg := addr(),addr()
	fmt.Printf("%T,%T\n",pos,neg)
	for i := 0; i<10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}

	f := fibonacci()
	for i := 0;i<10;i++ {
		// fmt.Println(fibonacci2(i))
		fmt.Println(f())
	}
}