package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {

	out, err := exec.Command("git", "clone", "https://github.com/Aswath-RA/Helloworld.git").Output()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(out))
}
