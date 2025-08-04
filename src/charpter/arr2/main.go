package main

import (
	"fmt"
	"sort"
)

func main() {
	var arr = [5]int{2, 33, 44, 29, 31}
	sort.Ints(arr[:])
	var n int
	fmt.Scanln(&n)
	var arr_cp = arr[:]
	arr_cp = append(arr_cp, n)
	sort.Ints(arr_cp)
	var arr_new [6]int
	copy(arr_new[:], arr_cp)
	fmt.Println(arr_new)
}
