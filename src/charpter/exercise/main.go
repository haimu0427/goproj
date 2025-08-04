package main

import (
	"fmt"
)

func peach(n int) int {
	if n == 10 {
		return 1
	} else {
		return (peach(n+1) + 1) * 2
	}
}
func main() {
	fmt.Println("peach(1) =", peach(1))
}
