package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	_ "strings"
)

func main() {
	fmt.Println("Welcome to Golang")
	PerformPostFormRequest()

}
func PerformPostFormRequest() {
	const myurl = "http://localhost/"

	//formdata

	data := url.Values{}
	data.Add("firstname", "Aswath")
	data.Add("lastname", "RA")
	data.Add("email", "aswath.ra@gmail.com")

	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))

}
