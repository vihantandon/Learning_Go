package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in go")

	myCh := make(chan int, 2)
	wg := &sync.WaitGroup{}
	// myCh <- 5
	// fmt.Println(<-myCh)
	// data can only be entered in the channel if someone is listening to it
	wg.Add(2)
	go func(ch <-chan int, wg *sync.WaitGroup) { //receive only . So not allowed to put up a closed statement inside it
		val, open := <-ch
		fmt.Println(open)
		fmt.Println(val) //one listener
		fmt.Println(val) //second listener
		wg.Done()
	}(myCh, wg)
	go func(ch chan<- int, wg *sync.WaitGroup) { //send only
		ch <- 5
		ch <- 9
		close(ch)
		wg.Done()
	}(myCh, wg)

	wg.Wait()
}
