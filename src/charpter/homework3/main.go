package main

import (
	"fmt"
)

func main() {
	var count int = 0
	for i := 2; i <= 100; i++ {

		if isPrime(i) {
			fmt.Printf("%-3d ", i)
			count++
			if count%5 == 0 {
				fmt.Println()
			}
		}
	}

}
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
