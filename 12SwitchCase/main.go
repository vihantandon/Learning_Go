package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch and case in go")

	rand.Seed(time.Now().UnixNano())
	dicenum := rand.Intn(6) + 1
	fmt.Println("Value of dice is: ", dicenum)

	switch dicenum {
	case 1:
		fmt.Println("you got 1")
	case 2:
		fmt.Println("you got 2")
	case 3:
		fmt.Println("you got 3")
	case 4:
		fmt.Println("you got 4")
		fallthrough // this will execute the next case as well
	case 5:
		fmt.Println("you got 5")
	case 6:
		fmt.Println("you got 6")
	default:
		fmt.Println("invalid dice number")
	}
}
