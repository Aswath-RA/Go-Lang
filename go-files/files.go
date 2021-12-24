package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to Go Files")
	text := "This is golang tutorial for files"

	file, err := os.Create("/home/blackfox/Desktop/gotutor/go-files/mytext.txt")

	if err != nil {
		panic(err)
	}
	length, err := io.WriteString(file, text)
	if err != nil {
		panic(err)
	}
	fmt.Println("Length of the String : ", length)

	readfile("/home/blackfox/Desktop/gotutor/go-files/mytext.txt")

	defer file.Close()

}
func readfile(filename string) {

	data, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}
	fmt.Println("You text in that file is: \n", string(data))

}
