package main

import (
	"fmt"
)

func test2(n1 int, n2 int) (int, int) {
	res1 := n1 + n2
	res2 := n1 - n2
	return res1, res2
}

func main() {
	fmt.Println("Hello, World!")
	res1, _ := test2(10, 5)
	fmt.Printf("Result%vniaho ", res1)
}
