package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func CopyFile(src, dst string) (written int64, err error) {
	file, err := os.Open(src)
	if err != nil {
		fmt.Printf("open file error:%v", err)
		return 0, err
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	dstFile, err := os.OpenFile(dst, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open file error:%v", err)
		return 0, err
	}
	writer := bufio.NewWriter(dstFile)
	defer dstFile.Close()
	return io.Copy(writer, reader)
}
func main() {
	_, err := CopyFile("玉子.jpg", "玉子_copy.jpg")
	if err != nil {
		fmt.Printf("copy file error:%v", err)
	} else {
		fmt.Println("File copied successfully")
	}
}
