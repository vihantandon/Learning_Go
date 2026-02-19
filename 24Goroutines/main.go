package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}
var wg sync.WaitGroup //pointers

var mut sync.Mutex //pointers

func main() {
	// go greeter("hello")
	// greeter("world")

	websitelist := []string{
		"https://google.com",
		"https://go.dev",
		"https://github.dev",
	}

	for _, web := range websitelist {
		wg.Add(1)
		go getStatusCode(web)
	}
	wg.Wait()
	fmt.Println(signals)
}

// func greeter(s string) {
// 	for i := 0; i < 5; i++ {
// 		time.Sleep(3 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)
	if err != nil {
		panic(err)
	} else {
		mut.Lock()
		// defer mut.Unlock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Printf("%d status code for %s \n", res.StatusCode, endpoint)
	}
}
