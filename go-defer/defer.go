package main

import "fmt"

func main() {
	var name string
	fmt.Print("Enter a name to reverse: ")
	fmt.Scan(&name)

	for i := 0; i < len(name); i++ {
		defer fmt.Printf(" Reversed name: %c", name[i])

	}

}
