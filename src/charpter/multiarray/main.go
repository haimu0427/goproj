package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	var arr [10]int
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(100)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(arr[:])))
	fmt.Println("reversed array:", arr)
	var sum float64
	for _, v := range arr {
		sum += float64(v)
	}
	average := sum / float64(len(arr))
	fmt.Println("average:", average)
	index, found := sort.Find(len(arr), func(i int) int {
		return 55 - arr[i]
	})
	if found {
		fmt.Println("index of 55:", index)
	} else {
		fmt.Println("55 not found in the array")
	}
}
