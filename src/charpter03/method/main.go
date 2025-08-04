package main

import (
	"fmt"
)

func zhuanzhi(dioarray *[3][3]int) {

	for i := 0; i < 3; i++ {
		for j := i + 1; j < 3; j++ {
			dioarray[i][j], dioarray[j][i] = dioarray[j][i], dioarray[i][j]
		}
	}
}

func main() {
	var dioarray = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	dio := dioarray[:][:]
	fmt.Println("Original Slice:")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%d ", dio[i][j])
		}
		fmt.Println()
	}
	fmt.Println("Original Array:")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%d ", dioarray[i][j])
		}
		fmt.Println()
	}

	zhuanzhi(&dioarray)
	fmt.Println("Transformed Array:")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%d ", dioarray[i][j])
		}
		fmt.Println()
	}
}
