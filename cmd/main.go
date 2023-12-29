package main

import (
	"fmt"
	"strconv"
	"time"
)

func unixTimeStampParser(timeStamp string) {
	numbers, err := strconv.ParseInt(timeStamp, 10, 32)
	if err != nil {
		panic("Invalid Timestamp")
	}
	readableTime := time.Unix(numbers, 0)
	fmt.Println(readableTime)
}

func main() {
	unixTimeStampParser("1703890777")
}
