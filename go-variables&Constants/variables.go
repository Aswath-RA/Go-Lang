package main
 
 import "fmt"
 
 var (
 	number int = 12345
 	age = 21
 	education = "BCA"
 )
 
 func main() {
 	fmt.Println("*** This Program is all about go Variables *** \n")
 	
 	/*  To declare a variable go Syntax is
 	  1)  var variable-name datatype = value
 	  2)  var variable-name = value
 	  3)  variable-name := value
 	*/
 	
 	var id int = 101
 	var name = "Aswath"
 	city := "Trichy"
 	
 	
 	// Printing the Above values
 	
 	fmt.Println("Your Name is      : ",name)
	fmt.Println("Your Id is        : ",id) 
	fmt.Println("Your Age is       : ",age)
	fmt.Println("Your Education is : ",education)
	fmt.Println("Your Number is    : ",number)
	fmt.Println("Your City is      : ",city)
	
	fmt.Println("\n *********************")
	
 	
 	
 	 
 }