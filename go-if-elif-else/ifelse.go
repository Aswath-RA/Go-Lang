package main 

import "fmt"

func main() {

	//simple if 
	x:=23
	fmt.Println("Value of X is : ",x)
	
	if x>15 {
		fmt.Println("X is greater than 15")
	}
	
	//simple if-else
	
	if x>23 {
		fmt.Println("X is greater than 23")
	} else {
		fmt.Println("X is less than 23")
	}
	
	//simple if-else if -else
	
	if x >23 {
		fmt.Println("X is greater than 23")
	} else if x>17 {
		fmt.Println("X is greater than 17")
	} else {
		fmt.Println("X is less than 23 and 17")
	}
	
	//simple nested-if
	
	if x>15 {
		if x>25{
		 	fmt.Println("X is greater than 25")
		 	}else {
		 	fmt.Println("x is greater than 15 but not greater than 25")
		 	}
	} else {
		fmt.Println("X is not greater than 15")
	}
}