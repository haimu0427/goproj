package main

import (
	"fmt"
)

type Cat struct {
	Name  string
	Age   int
	Color string
	Hobby string
}

func (c Cat) Meow() {
	fmt.Println(c.Name, "says Meow!")
}
func (c *Cat) String() string {
	str := fmt.Sprintf("Cat Name: %s, Age: %d, Color: %s, Hobby: %s", c.Name, c.Age, c.Color, c.Hobby)
	return str
}

func main() {
	cat := Cat{
		Name:  "Whiskers",
		Age:   3,
		Color: "Gray",
		Hobby: "Chasing mice",
	}
	cat.Meow()
	fmt.Println(cat.String())
}
