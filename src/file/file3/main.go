package main

import (
	"fmt"
	"os"
)

func main() {
	content, err := os.ReadFile("file1.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	file2, err := os.OpenFile("file2.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file2.Close()
	file2.Write(content)
}
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
