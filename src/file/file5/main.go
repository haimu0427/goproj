package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type CharCount struct {
	ChCount    int
	NumCount   int
	SpaceCount int
	OtherCount int
}

func main() {
	file, err := os.Open("abc.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var count CharCount
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		rune1 := []rune(str)
		for _, v := range rune1 {
			switch {
			case v >= 'a' && v <= 'z', v >= 'A' && v <= 'Z':
				count.ChCount++
			case v >= '0' && v <= '9':
				count.NumCount++
			case v == ' ' || v == '\t':
				count.SpaceCount++
			default:
				count.OtherCount++
			}
		}
		if err == io.EOF {
			break
		}
	}
	fmt.Printf("Character Count: %d\nNumber Count: %d\nSpace Count: %d\nOther Count: %d\n", count.ChCount, count.NumCount, count.SpaceCount, count.OtherCount)
}
