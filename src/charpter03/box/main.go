package main

import (
	"fmt"
)

type Box struct {
	Width  int
	Height int
	Depth  int
}

func (b Box) GetVolume() int {
	return b.Width * b.Height * b.Depth
}
func main() {
	box := Box{}
	fmt.Println("Enter the width, height, and depth of the box:")
	fmt.Scan(&box.Width, &box.Height, &box.Depth)
	fmt.Printf("The volume of the box is: %d\n", box.GetVolume())
}
