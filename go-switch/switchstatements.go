package main

import "fmt"

func main() {

 	var word string
	fmt.Print("Check your character is a vowel : ")
 	fmt.Scan(&word)
 	
 		
 	switch word {
 	case "a","e","i","o","u","A","E","I","O","U":
 		fmt.Print(" Your char is vowel")
 	default:
 		fmt.Println("Yout char is not a vowel")
 		}
}