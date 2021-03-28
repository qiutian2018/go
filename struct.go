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

func (hero *Hero) setName(s string) string {
	hero.Name = s
	return hero.Name
}

func GetName(hero Hero) string {
	return hero.Name
}

func GetName2(hero *Hero) string {
	return hero.Name
}

type MyInt int
func (i MyInt) test() int {
	return int(i)
}

func SetName(hero *Hero,s string) string {
	hero.Name = s
	return hero.Name
}

func main(){
	people1 := Hero{"张三",22}
	people2 := &people1
	fmt.Println(people1.getName())
	fmt.Println(people2.getName())
	fmt.Println((&people1).getName())
	fmt.Println(people1.setName("张三别名"))
	fmt.Println(people1.getName())
	fmt.Println(GetName(people1))

	i := MyInt(2)
	fmt.Println(i.test())

	SetName(&people1,"张三别名2")
	fmt.Println(people1.getName())
	fmt.Println(GetName(people1))
	fmt.Println(GetName(*people2))

	people3 := &Hero{"李四",23}
	fmt.Println(GetName2(people3))
}