package main
import (
	"fmt"
)

type Animal interface {
	Say() string
	M()
}

type I interface {
	M()
}

type Dog string
type Cat string
type Unknow string

func (dog Dog) getName() string {
	return string(dog)
}
func (cat Cat) getName() string {
	return string(cat)
}
func (dog Dog) Say() string {
	return dog.getName() +":"+ "汪汪汪~"
}
func (dog Dog) M() {
	fmt.Println(dog.getName() + " M()")
}
func (cat *Cat) Say() string {
	return cat.getName() +":"+ "喵喵喵~"
}
func (cat *Cat) M(){
	fmt.Println(cat.getName() + " M()")
}
func (unknow Unknow) Say() string {
	return "呜呜呜~"
}
func (unknow Unknow) M(){
	fmt.Println("M()")
	return
}

func CheckType(i interface {}) {
	switch t := i.(type) {
		case string:
			fmt.Printf("%v is string\n",t)
		case int:
			fmt.Printf("%v is int\n",t)
		case float64:
			fmt.Printf("%v is float64\n",t)
		default:
			fmt.Printf("%v is unknow\n",t)
	}
}
func main() {
	var animal,animal2,animal3 Animal
	dog := Dog("小狗")
	cat := Cat("小猫")
	wu := Unknow("未知动物")
	animal = dog
	animal2 = &cat
	animal3 = wu
	fmt.Println(animal.Say(),animal2.Say(),animal3.Say())
	animal.M()
	animal2.M()
	animal3.M()

	// var i I
	// i = dog
	// i.M()
	var i2,i3 interface {}
	i2 = 11
	
	fmt.Printf("%v,%T\n",i2,i2)
	fmt.Printf("%v,%T,%v\n",i3,i3,i3==nil)

	var ii interface {} = "hello go"
	// s,_ := ii.(string)
	s,_ := ii.(int)
	fmt.Println(s)
	fmt.Printf("%T\n",s)

	CheckType("aaa")
	CheckType(2)
}