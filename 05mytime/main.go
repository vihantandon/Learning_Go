package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Time Study")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday")) //Fixed date and time for formatting
	fmt.Println(presentTime.Month())

	createTime := time.Date(2020, time.March, 10, 23, 56, 34, 0, time.UTC)
	fmt.Println(createTime)
	fmt.Println(createTime.Day())
}
