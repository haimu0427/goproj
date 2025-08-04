package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	// get the current time
	currentTime := time.Now()
	fmt.Println("Current Time:", currentTime)
	str := ""
	for i := 0; i < 10; i++ {
		str = str + strconv.Itoa(i)
	}
	latestTime := time.Now()
	fmt.Println("Latest Time:", latestTime)
	fmt.Println("Time Difference:", latestTime.Sub(currentTime))
}
