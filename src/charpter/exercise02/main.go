package main

import (
	"fmt"
)

var age = 25

func init() {
	fmt.Println("init function called")
}

func AddUpper() func(int) int {
	var n int = 10
	var str string = "hello"
	return func(x int) int {
		n += x
		str += fmt.Sprint(x)
		fmt.Println("Current value of n:", n)
		fmt.Println("Current value of str:", str)
		return n
	}
}

func main() {
	fmt.Println("Hello, World!")
	res1 := func(a int, b int) int {
		return a + b
	}(10, 20)
	fmt.Println("Result:", res1)

	fun := func(a int, b int) int {
		return a * b
	}
	res2 := fun(10, 20)
	fmt.Println("Result:", res2)

	addUpper := AddUpper()
	fmt.Println(addUpper(1))
	fmt.Println(addUpper(2))
	fmt.Println(addUpper(3))
}
