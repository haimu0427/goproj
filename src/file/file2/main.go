package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.OpenFile("nihao.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		fmt.Print(line)
		if err == io.EOF {
			break
		}
	}

	for i := range 5 {
		fmt.Fprintf(writer, "我当二十不得意，\n一心愁谢如枯兰。\n")
		fmt.Fprintf(writer, "%d", i+1)
	}
	writer.Flush()

	file1, err := os.OpenFile("test.txt", os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file1.Close()
	write := bufio.NewWriter(file1)
	write.WriteString("坤坤\n")
	write.Flush()
}
