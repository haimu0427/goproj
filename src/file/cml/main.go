package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args)
	fmt.Println(len(os.Args))
	for _, v := range os.Args {
		fmt.Println(v)
	}
}
