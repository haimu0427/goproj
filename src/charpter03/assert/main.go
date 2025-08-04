package main

import (
	"fmt"
)

type Point struct {
	X int
	Y int
}

func main() {
	var a any
	var p Point
	p.X = 10
	p.Y = 20
	a = p
	fmt.Println("Point:", a)

	var b Point
	fmt.Println("Point after assertion:", b)

	if b, ok := a.(Point); ok {
		fmt.Println("Point after assertion:", b)
	} else {
		fmt.Println("a is not of type Point")
	}
	fmt.Println("继续进行")

}
