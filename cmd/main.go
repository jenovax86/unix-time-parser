package main

import (
	"fmt"
	"strconv"
	"time"
)

func unixTimeStampParser() {
	var timeStamp string
	fmt.Scan(&timeStamp)
	numbers, err := strconv.ParseInt(timeStamp, 10, 32)
	if err != nil {
		panic("Invalid Timestamp")
	}
	readableTime := time.Unix(numbers, 0)
	fmt.Println(readableTime)
}

func main() {
	unixTimeStampParser()
}
