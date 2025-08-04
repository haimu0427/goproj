package main

import (
	"fmt"
)

type Visitor struct {
	Name string
	Age  int
}

func (v Visitor) AreYouFree() bool {
	return v.Age < 18
}
func main() {
	visitor := Visitor{}
	v2 := Visitor{"John Doe", 20}
	fmt.Println(v2)
	fmt.Println("Enter your name and age:")
	fmt.Scan(&visitor.Name, &visitor.Age)
	if visitor.AreYouFree() {
		fmt.Printf("Hello %s, you are free to enter.\n", visitor.Name)
	} else {
		fmt.Printf("Sorry %s, you are not allowed to enter.\n", visitor.Name)
	}
}
