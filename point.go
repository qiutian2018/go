package main
import (
	"fmt"
)

type S1 struct {
	x int
	y int
}

func main(){
	i,j := 1,2
	p := &i
	*p = 3
	fmt.Println(p,*p,&i,j)
	p = &j
	fmt.Println(*p,j)
	*p = *p/2
	fmt.Println(*p,j)

	s := S1{1,2}
	s.y = 4
	fmt.Println(s.x,s.y,s)
	p2 := &s
	p2.x = 1e9
	fmt.Println(s,p2,*p2)

	v1 := S1{2,3}
	v2 := S1{y:3}
	v3 := S1{}
	v4 := &S1{4,4}
	fmt.Println(v1,v2,v3,v4,*v4)
}