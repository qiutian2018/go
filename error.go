package main
import (
	"fmt"
)

type MyError struct {
	When int
	What string
}

func GetType(i interface{}) string{
	switch t := i.(type) {
		case string:
			fmt.Println("%v\n",t)
			return "string";
		case int:
			fmt.Println("%v\n",t)
			return "int"
		default:
			fmt.Println("%v\n",t)
			return "other"
	}
}

func (err *MyError) Error() string {
	if t := err.When; t < 1 {
		return fmt.Sprintf("%v is small than 1\n",err.When)
	}
	return ""
}

func run() error {
	return &MyError{
		0,
		"zhangsan",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
	// num := 10
	// GetType(num)
	// fmt.Println(GetType(num)=="string")
}