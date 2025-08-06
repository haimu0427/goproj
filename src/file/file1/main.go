package main

import (
	"bufio"
	"fmt"
	_ "io"
	"os"
)

func main() {
	file, err := os.Open("text.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		fmt.Printf("%s", string(line))
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			fmt.Println(err)
			return
		}
	}
	fmt.Printf("我读完了, 夸我一下\n")

	content, err := os.ReadFile("text.txt")
	if err != nil {
		fmt.Println("error is :", err)
		return
	}
	fmt.Println(string(content))

	// scanner := bufio.NewScanner(file)
	// for scanner.Scan() {
	// 	line := scanner.Text()
	// 	fmt.Println(line)
	// }
	// if content, err := io.ReadAll(file); err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(string(content))
	// }

	defer file.Close()
}
