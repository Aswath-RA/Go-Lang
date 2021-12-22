package main

import "fmt"

func main() {

	//Operators
	
	/* Types of operator in go
	   1) Arithmetic 
	   2) Assignment 
	   3) Comparison 
	   4) Logical
	   5) Bitwise
	 */
	 
	 
	 //Arithmetic operator
	 
	 a:=100
	 b:=20
	 fmt.Println("Arithmetic Operator \n")
	 fmt.Println("Addition of two variables       : ", a+b)
	 fmt.Println("Subtraction of two variables    : ", a-b)
	 fmt.Println("Multiplication of two variables : ", a*b)
	 fmt.Println("Divison of two variables        : ", a/b)
	 fmt.Println("Modulus of two variables        : ", a%b)
	 a++
	 fmt.Println("Increment of two variables      : ", a)
	 a--
	 fmt.Println("Decrement of two variables      : ", a)
	 fmt.Println("")
	 
	 //Assignment operator
	 
	 var x=5
	 fmt.Println("Assignment Operator \n")
	 fmt.Println("Value of X is     :",x)
	 x+=5
	 fmt.Println("Value of X=+5 is  :",x)
	 x-=5
	 fmt.Println("Value of X=-5 is  :",x)
	 x*=2
	 fmt.Println("Value of X=*2 is  :",x)
	 x/=2
	 fmt.Println("Value of X=/2 is  :",x)
	 x%=2
	 fmt.Println("Value of X=%2 is  :",x) 
	 fmt.Println("")
	 
	 //Comparison operator
	 
	 var m=5
	 var n=7
	 
	 fmt.Println("Comparison Operator \n")
	 fmt.Println("Value of M                       : ",m)
	 fmt.Println("Value of N                       : ",n)
	 fmt.Println("M is  Equal to n                 : ",m==n)
	 fmt.Println("M is  Not Equal to n             : ",m!=n)
	 fmt.Println("M is Greater than or  Equal to n : ",m>=n)
	 fmt.Println("M is Less than or  Equal to n    : ",m<=n)
	 fmt.Println("M is Greater than N              : ",m>n)
	 fmt.Println("M is Less than N                 : ",m>n)
	 fmt.Println("\n")
	 
	 //Logical Operator
	 fmt.Println("Logical Operator \n")
	 fmt.Println("And operator m>3 && n>3 : ",m>3 && n>3 )
         fmt.Println("And operator m>3 || n>3 : ",m>6 || n>3 )
         fmt.Println("And operator m>3 ! n>3 : ",!(m>4 && n>5) )

	//Bitwise Operator
	fmt.Println("Bitwise operator \n")
	fmt.Println(" Bitwise &   m & n : ",m & n)
	fmt.Println(" Bitwise |   m | n : ",m | n)
	fmt.Println(" Bitwise ^   m ^ n : ",m ^ n)
	 
	  
}