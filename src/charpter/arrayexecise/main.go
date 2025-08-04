package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var chararr [26]byte
	for i, _ := range chararr {
		chararr[i] = 'a' + byte(i)
	}
	fmt.Printf("chararr: %c\n", chararr)
	var arr1 = [5]int{22, 22, 123, 4, 5}
	var average float64
	var sum int
	for _, v := range arr1 {
		sum += v

	}
	average = float64(sum) / float64(len(arr1))
	fmt.Printf("average: %f\n", average)

	var intarr3 [5]int
	for i := 0; i < len(intarr3); i++ {
		intarr3[i] = rand.Intn(100)
	}
	fmt.Println("intarr3:", intarr3)

	//reverse
	for i := 0; i < len(intarr3)/2; i++ {
		intarr3[i], intarr3[len(intarr3)-i-1] = intarr3[len(intarr3)-i-1], intarr3[i]
	}
	fmt.Println("reversed intarr3:", intarr3)
}
