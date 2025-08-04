package main

import (
	"fmt"
)

func main() {
	var arr [4][4]int
	for i := 0; i < len(arr)*len(arr[0]); i++ {
		fmt.Scan(&arr[i/4][i%4])
	}
	for _, v := range arr {
		fmt.Println(v)
	}

	fmt.Println("交换后的数组")
	arr[1], arr[2] = arr[2], arr[1]
	arr[0], arr[3] = arr[3], arr[0]
	for i := range arr {
		fmt.Println(arr[i])
	}
}
