package main

import (
	"fmt"
	"os"
)

func main() {
	var s string
	fmt.Println(len(os.Args))
	for i := 0; i < len(os.Args); i++ {
		s += os.Args[i] + " "
	}
	fmt.Println(s)
}
