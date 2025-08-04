package main

import (
	"fmt"
)

func main() {
	var arr [3][4]int
	for i := 0; i < len(arr)*len(arr[0]); i++ {
		fmt.Scan(&arr[i/4][i%4])
	}
	for _, v := range arr {
		fmt.Println(v)
	}
	fmt.Println("清空周围后的数组:")
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			if i == 0 || i == len(arr)-1 || j == 0 || j == len(arr[i])-1 {
				arr[i][j] = 0
			}
		}
	}
	for _, v := range arr {
		fmt.Println(v)
	}
}
