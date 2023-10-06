package main

import "fmt"

type action interface {
	say(msg string)
}

type person struct {
	name string
	age  int
}

func (person) say(msg string) {

}
func getname(p person) string {
	return p.name
}
func main() {
	var p = person{name: "bit", age: 23}
	fmt.Printf("the person age:%d\n", p.age)
	fmt.Println(getname(p))
}
