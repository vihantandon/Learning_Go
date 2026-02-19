package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Get request in go")
	// PerformGetRequest()
	// PerformPostRequest()
	PerformPostFormRequest()
}

func PerformGetRequest() {
	const myurl = "http://localhost:8000/get"

	res, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	fmt.Println("Status code : ", res.StatusCode)
	fmt.Println("Content Length : ", res.ContentLength)

	content, _ := ioutil.ReadAll(res.Body)
	// fmt.Println(string(content)) method 1 to convert byte array to string

	var responseString strings.Builder
	byteCount, _ := responseString.Write(content)
	fmt.Println("bytecount is : ", byteCount)
	fmt.Println("responseString is : ", responseString.String()) // method 2 to convert byte array to string

}

func PerformPostRequest() {
	const myurl = "http://localhost:8000/post"

	//json data
	requestBody := strings.NewReader(` 
		{
			"name" : "vihan",
			"age" : "21",
			"city" : "Rewa"

		}
	`)

	res, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	content, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(content))
}

func PerformPostFormRequest() {
	const myurl = "http://localhost:8000/postform"

	//formdata

	data := url.Values{}
	data.Add("name", "vihan")
	data.Add("age", "21")
	data.Add("city", "Rewa")

	res, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	content, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(content))
}
