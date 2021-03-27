package main
import (
	"fmt"
	"math"
)

func sqrt(x float64) string{
	if x<0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x,n,lim float64) float64{
	if v := math.Pow(x,n); v < lim{
		fmt.Println("---",x,n,v,lim,"---")
		return v
	}else{
		fmt.Println("v is ok",x,n,v,lim,"\n")
	}
	fmt.Println("***",x,n,lim,"****")
	return lim
}

func Sqrt(x float64) float64{
	x -= (x*x - x) / (2*x)
	return x
}

func main(){
	fmt.Println("hello go")
	sum := 1
	for ; sum<30; {
		sum += sum
	}
	fmt.Println(sum)


	sum2 := 1
	for sum2 < 10 {
		sum2 += sum2
	}
	fmt.Println(sum2)

	i := 1
	for i<10 {
		i += i
		fmt.Println("for",i)
	}

	if 2>1 {
		fmt.Println("2>1")
	}

	fmt.Println(sqrt(2),sqrt(-4))

	fmt.Println(
		pow(2,2,4),
		pow(3,2,10),
		pow(3,3,20),
	)

	fmt.Println("------Sqrt------\n")
	fmt.Println(Sqrt(2))
}