package main
import (
	"fmt"
	"runtime"
	"time"
)

func macOs() string{
	return "darwin"
}
func printString(s string){
	fmt.Println(s)
}
func main(){
	fmt.Println("Go runs on")
	switch os := runtime.GOOS; os {
		case macOs():
			fmt.Println("OS X")
		case "linux":
			fmt.Println("Linux")
		default:
			fmt.Println("Windows")
	}

	switch v := 1.3; v {
		case 1.3:
			fmt.Println("=1.3")
		case 2.3:
			fmt.Println("=2.3")
		default:
			fmt.Println("unknow value")
	}

	fmt.Println(time.Now().Weekday())

	defer fmt.Println("two")
	defer fmt.Println("three")
	fmt.Println("one")
	defer printString("four")
	defer printString("five")
}