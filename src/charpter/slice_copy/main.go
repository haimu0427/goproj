package main

import (
	"fmt"
)

func main() {
	str := "hello@world"
	slice := str[5:]
	fmt.Println("Slice:", slice)
	slice_str := str[:]
	fmt.Println("Slice_str:", slice_str)
	byteSlice := []byte(slice_str)
	byteSlice[0] = 'S'
	modifiedSliceStr := string(byteSlice)
	fmt.Println("Modified Slice_str:", modifiedSliceStr)
	runeSlice := []rune(slice_str)
	runeSlice[0] = '中'
	modifiedSliceStr = string(runeSlice)
	fmt.Println("Modified Slice_str with rune:", modifiedSliceStr)
}
