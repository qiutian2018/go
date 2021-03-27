package main
import (
	"fmt"
)

type Hero struct {
	Name string
	Age int
}

func (hero Hero) getName() string {
	return hero.Name
}

// todo::
func (hero Hero) setName(s string) string {
	hero.Name = s
	return hero.Name
}

func GetName(hero Hero) string {
	return hero.Name
}

type MyInt int
func (i MyInt) test() int {
	return int(i)
}

func main(){
	people1 := Hero{"张三",22}
	fmt.Println(people1.getName())
	fmt.Println(people1.setName("张三别名"))
	fmt.Println(people1.getName())
	fmt.Println(GetName(people1))

	i := MyInt(2)
	fmt.Println(i.test())
}