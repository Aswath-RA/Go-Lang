package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	fmt.Println("File.io in golang")

	file, err := os.Create("./mytxt.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("File Created")

	content := "Hi,This is Golang file.io"

	length, err := io.WriteString(file, content)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("File length: ", length)
	defer file.Close()

	data, err := ioutil.ReadFile("./mytxt.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))

}
