package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://aswath.com:6060/about?jobrole=SE&age=21"

func main() {
	fmt.Println("Welcome to handling URLs in golang")
	fmt.Println("URL: ", myurl)
	fmt.Println()
	//parsing
	result, _ := url.Parse(myurl)

	fmt.Println("Scheme : ", result.Scheme)
	fmt.Println("Host   : ", result.Host)
	fmt.Println("Path   : ", result.Path)
	fmt.Println("Result : ", result.Port())
	fmt.Println("Query  : ", result.RawQuery)

	qparams := result.Query()

	fmt.Printf("The type of query params are: %T\n", qparams)
	fmt.Println()

	fmt.Println(qparams["jobrole"])
	fmt.Println()

	for _, val := range qparams {
		fmt.Println("Param is: ", val)
	}
	fmt.Println()
	//Url building .

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "aswath.com",
		Path:    "/about",
		RawPath: "?user=aswath",
	}

	anotherURL := partsOfUrl.String()
	fmt.Println("Url builded: ", anotherURL)

}
