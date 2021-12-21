package main 

import "fmt"

func main() {
	
	//Data types in Go.
	
	//Boolean
	
	var limit_reached bool = true
	number := true
	
	//int
	
	var age int =21
	height := 174
	activity:= -1
	
	//Float
	
	var run_speed  float32 =35.6
	weight := 57.9 
	
	//String 
	
	var name string = "rock"
	edu :="BE"
	
	fmt.Println("Limit Reached   : ", limit_reached)
	fmt.Println("Number is there : ", number)
	fmt.Println("Age is          : ", age)
	fmt.Println("Height is       : ", height)
	fmt.Println("Activity is     : ", activity)
	fmt.Println("Run Speed is    : ", run_speed)
	fmt.Println("Weight is       : ", weight)
	fmt.Println("Name is         : ", name)
	fmt.Println("Education is    : ", edu)
}