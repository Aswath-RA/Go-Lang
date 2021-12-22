package main 

import "fmt"

func main() {
	message()
	add(5,5)
	fmt.Println("Multiplication of two numbers: ",multiply(4,4))
	fmt.Println("Factorial of given number 4: ",fact(4))
}

//simple function
func message() {
	fmt.Println("I am learning Go-lang ")
}

//parameter passing 

func add(x int ,y int) {
	fmt.Println("Addition of two numbers: ",x+y)
}

//returning value
func multiply(m int ,n int) int {
	result :=m*n
	return result
}

//Recursion function


func fact(x int )(y int) {
	if x>0{
		y=x*fact(x-1)
    }else {
    	y=1
    }
        
    return 
}