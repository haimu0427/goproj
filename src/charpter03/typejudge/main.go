package main

import (
	"fmt"
)

type Student struct {
	Name string
	Age  int
}

func TypeJudge(items ...any) {
	for index, item := range items {
		switch item.(type) {
		case int, int8, int16, int32, int64:
			fmt.Printf("Item %d is of type int: %d\n", index, item)
		case string:
			fmt.Printf("Item %d is of type string: %s\n", index, item)
		case bool:
			fmt.Printf("Item %d is of type bool: %t\n", index, item)
		case float64:
			fmt.Printf("Item %d is of type float64: %f\n", index, item)
		case Student:
			fmt.Printf("Item %d is of type Student: %+v\n", index, item)
		case *Student:
			fmt.Printf("Item %d is of type *Student: %+v\n", index, *(item).(*Student))
		default:
			fmt.Println("Item", index, "is of unknown type:", item)
		}
	}
}
func main() {
	var a any
	a = 42
	b := "Hello"
	c := true
	d := 3.14
	var s Student
	s.Name = "Alice"
	s.Age = 20
	var p *Student
	p = &s

	TypeJudge(a, b, c, d, s, p)

	fmt.Println("继续进行")
}
