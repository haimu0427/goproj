package main

import (
	"fmt"
)

func getSum(n1 int, n2 int) int {
	return n1 + n2
}

func myfun(funcvar func(int, int) int, num1 int, num2 int) int {
	return funcvar(num1, num2)
}

func sum(n1 int, args ...int) int {
	sum := n1
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}

func main() {
	a := getSum
	fmt.Printf("a:%T\n", a)

	result := myfun(getSum, 50, 60)
	fmt.Println("result:", result)

	type myInt int
	var c int = 10
	var b myInt = 20
	fmt.Printf("c:%T, b:%T\n", c, b)
	fmt.Println("c:", c, "b:", b)
	c = int(b)

}
