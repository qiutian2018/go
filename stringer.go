package main
import (
	"fmt"
)

type Hero struct {
	Name string
	Age int
}

func (hero Hero) String() string {
	return fmt.Sprintf("%v (%v years old)\n",hero.Name,hero.Age)
}

type IPAddr [4]byte

func (ip IPAddr) String() string {
	var tmp string
	for _,v := range ip {
		tmp += fmt.Sprintf("%v.",v)
	}
	return tmp
}

func main(){
	hero1 := Hero{"张三",22}
	hero2 := Hero{"李四",23}
	fmt.Println(hero1,hero2)

	ip1 := map[string]IPAddr{
		"aaa":{1,2,3,4},
		"bbb":{11,21,31,41},
	}
	fmt.Println(ip1)
	for name,ip := range ip1 {
		fmt.Printf("%v:%v\n",name,ip)
	}
}