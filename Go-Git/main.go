package main

import (
	"log"
	"os"

	"gopkg.in/src-d/go-git.v4"
)

func main() {
	_, err := git.PlainClone("/home/blackfox/Desktop/gotutor/Go-Git/output", false, &git.CloneOptions{
		URL:      "https://github.com/Aswath-RA/Helloworld.git",
		Progress: os.Stdout,
	})
	if err != nil {
		log.Fatal(err)

	}

}
