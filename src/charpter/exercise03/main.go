package main

import (
	"fmt"
	"strings"
)

func makeSuffix(suffix string) func(string) string {

	return func(name string) string {
		var result string
		if !strings.HasSuffix(name, suffix) {
			result = name + suffix
		} else {
			result = name
		}
		return result
	}

}
func main() {
	fun := makeSuffix(".png")
	fmt.Println("File name with suffix:", fun("image"))
}
