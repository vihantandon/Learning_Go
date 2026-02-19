package main

import (
	"fmt"
	"net/url"
)

const myurl = "https://lco.dev:3000/learn?course=reactjs&paymentid=ghbj456ghb"

func main() {
	fmt.Println("Handelling Url's in go")
	fmt.Println(myurl)

	//parsing url

	result, _ := url.Parse(myurl)
	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port())
	// fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("Type of query params: %T \n", qparams)
	fmt.Println(qparams)
	fmt.Println(qparams["course"])

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=hitesh",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println("Another url is:", anotherUrl)

}
