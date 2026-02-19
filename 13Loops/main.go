package main

import "fmt"

func main() {
	fmt.Println("Loops in go")
	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	for d := 0; d < len(days); d++ {
		fmt.Println(d+1, days[d])
	}

	for i := range days { // range returns the index of the slice
		fmt.Println(days[i])
	}

	for index, days := range days {
		fmt.Printf("%v : %v \n", index, days)
	}

	val := 1
	for val <= 10 {

		if val == 2 {
			goto label
		}

		if val == 5 {
			val++ //only continue will cause an infinite loop because val will always be 5, so we need to increment it before continue
			continue
		}

		fmt.Println(val)

		val++
		// ++val is not allowed in go, only val++ is allowed
	}

label:
	fmt.Println("hey you stopped at 2 nigga")
}
