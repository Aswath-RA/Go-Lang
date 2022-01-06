package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"name"` //for creating correct way of json
	Platform string `json:"platform"`
}

func main() {
	fmt.Println("Welcome to go lang json")

	EncodeJson()
}

func EncodeJson() {
	learning := []course{
		{"aswath", "golang"}, {"rock", "docker"},
	}

	//package this data as JSON data

	finalJson, err := json.MarshalIndent(learning, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)

}
