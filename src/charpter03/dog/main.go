package main

import (
	"fmt"
)

type Dog struct {
	Name   string
	Age    int
	Weight float64
}

func (d Dog) Say() string {
	str := fmt.Sprintf("Woof! My name is %s, I am %d years old and I weigh %.2f kg.\n", d.Name, d.Age, d.Weight)
	return str
}
func main() {
	dog := Dog{
		Name:   "Buddy",
		Age:    5,
		Weight: 20.5,
	}
	fmt.Println(dog.Say())
}
