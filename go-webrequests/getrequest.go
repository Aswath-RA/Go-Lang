package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	getrequest()

}
func getrequest() {
	const url = "http://localhost/"
	result, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer result.Body.Close()
	fmt.Println("Status code :", result.StatusCode)
	fmt.Println("content length: ", result.ContentLength)

	content, _ := ioutil.ReadAll(result.Body)
	fmt.Println(string(content))

}
