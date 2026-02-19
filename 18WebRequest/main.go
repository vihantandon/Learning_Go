package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://google.com"

func main() {
	fmt.Println("Web Requests in go")
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Type of response: %T \n", response)
	// fmt.Println("Response:", response)

	// response.Body.Close() //Callers responsibility to close the connection

	databytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	content := string(databytes)
	fmt.Println(content)

}
