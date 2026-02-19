package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices")

	var list = []string{"yo", "go", "lang"}
	fmt.Printf("type is %T \n", list)

	list = append(list, "is", "awsome")
	fmt.Println(list)

	list = append(list[1:3]) //ending range is non inclusive
	fmt.Println(list)

	score := make([]int, 4, 10) //length and capacity

	score[0] = 100
	score[1] = 90
	score[2] = 80
	score[3] = 70
	// score[4] = 69 error because out of range

	score = append(score, 69, 70) // this will work bcz it'll reallocate the memory
	sort.Ints((score))
	fmt.Println(score)
	fmt.Println(sort.IntsAreSorted(score))

	fmt.Println("remove a value from slices")

	var s = []string{"a", "b", "c", "d", "e"}
	index := 2
	s = append(s[:index], s[index+1:]...)
	fmt.Println(s)
}
