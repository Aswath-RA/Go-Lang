package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to Golang")
	PerformPostJsonRequest()

}

func PerformPostJsonRequest() {
	const myurl = "http://localhost/"

	//fake json payload

	requestBody := strings.NewReader(`
		{
			"coursename":"golang"
		}
	`)

	response, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}
