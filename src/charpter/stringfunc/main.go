package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	n, err := strconv.Atoi("1234")
	if err != nil {
		fmt.Println("Error converting string to int:", err)
	} else {
		fmt.Println("Conversion successful, number is:", n)
	}

	str := strconv.Itoa(1234)
	fmt.Printf("类型是%T 值是%s\n", str, str)

	var byteSlice = []byte("hello go")
	fmt.Printf("字符的值是%v, 字符的类型是%T\n", byteSlice, byteSlice)

	str2 := string([]byte{104, 101, 108, 108, 111})
	fmt.Printf("字符串是%s, 类型是%T\n", str2, str2)

	str3 := strconv.FormatInt(1234, 2)
	fmt.Print(str3, "\n")

	fmt.Printf("contains is %s\n", strconv.FormatBool(strings.Contains("golang", "ol")))
	fmt.Printf("number is %d\n", strings.Count("golang", "g"))

	fmt.Println("result is", "Abc" == "abc")
	fmt.Println("Index is", strings.Index("olang", "g"))
	fmt.Println("LastIndex is", strings.LastIndex("golang", "g"))
	fmt.Println("repeat is", strings.Repeat("go", 3))

	strings.Replace("gogogo出发咯", "go", "kakaka", -1)

	strarr := strings.Split("hello,world", ",")

	fmt.Printf("切割后的数组是%v, 类型是%T\n", strarr, strarr)

	fmt.Println("to lower is", strings.ToLower("HHHH ni pa SHi yig shabi foo dog "))
	fmt.Println("to upper is", strings.ToUpper("HHHH ni pa SHi yig shabi foo dog "))
	fmt.Println("trim is", strings.TrimSpace("   hello shabi    "))
	fmt.Println("trim left is", strings.TrimLeft(" ! !hello shabi   !", " !"))
	fmt.Println("trim right is", strings.TrimRight("   hello shabi    ", " "))
	fmt.Println("trim all is", strings.Trim("   hello shabi    ", " "))

	fmt.Println("HasPrefix is", strings.HasPrefix("gola ng", "go"))
	fmt.Println("HasSuffix is", strings.HasSuffix("golang", "ang"))

}
