package main
import (
	"fmt"
	"strings"
)

type Vertex struct{
	Lat,Lng float64
}

var m map[string]Vertex

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	for _,v := range strings.Fields(s) {
		// fmt.Println(x,v)
		m[v] += 1
	}
	fmt.Printf("fields in s are:%q",strings.Fields(s))
	return m
}

func main(){
	m = make(map[string]Vertex)
	fmt.Println(len(m))
	m["beijing"] = Vertex{
		233.112,3111.22,
	}
	m["shanghai"] = Vertex{
		1222.1,222.1111,
	}
	m = map[string]Vertex{
		"aaa":Vertex{1.22,2.22,},
		"bbb":{2.22,2.333},
	}
	fmt.Println(len(m))
	// delete(m,"aaa")
	fmt.Println(len(m))
	fmt.Println(m,m==nil,"\n",m["beijing"],m["beijing2"])
	elem,ok := m["bbbwc"]
	fmt.Println(elem,ok)


	fmt.Println(WordCount("asdfasxxd this is go this is"))
}