package main

import (
	"fmt"
)

func fbn(n int) []uint64 {
	fbnSlice := make([]uint64, n+2)
	fbnSlice[1] = 1
	fbnSlice[0] = 1
	for i := 2; i < n; i++ {
		fbnSlice[i] = fbnSlice[i-1] + fbnSlice[i-2]
	}
	return fbnSlice
}
func main() {
	var n int
	fmt.Scanln(&n)
	fbnSequence := fbn(n)
	fmt.Println("Fibonacci sequence:", fbnSequence)
}
