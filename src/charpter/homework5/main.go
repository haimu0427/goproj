package main

import (
	"fmt"
)

func main() {
	for i := 'a'; i <= 'z'; i++ {
		fmt.Printf("%c ", i)
	}
	fmt.Println() // Print a newline after the loop
	for i := 'Z'; i >= 'A'; i-- {
		fmt.Printf("%c ", i)
	}

}
