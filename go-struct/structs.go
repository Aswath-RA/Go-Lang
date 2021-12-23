package main

import "fmt"

//Declaring Struct

type Employee struct {
    name string
    id int
    number int
    age int
    sex string
  }

func main(){
    var emp1 Employee //Struct variables 
    var emp2 Employee
    
    emp1.name = "Aswath"
    emp1.id =123
    emp1.number=456
    emp1.age=21
    emp1.sex="male"
    
    fmt.Println("Employee Details :\n")
    
    fmt.Println("Employee Name   : ",emp1.name)
    fmt.Println("Employee Id     : ",emp1.id)
    fmt.Println("Employee Number : ",emp1.number)
    fmt.Println("Employee Age    : ",emp1.age)
    fmt.Println("Employee Sex    : ",emp1.sex)
    fmt.Println()
    
    
    emp2.name = "rock"
    emp2.id =456
    emp2.number=789
    emp2.age=24
    emp2.sex="male"
    
    fmt.Println("Employee Name   : ",emp2.name)
    fmt.Println("Employee Id     : ",emp2.id)
    fmt.Println("Employee Number : ",emp2.number)
    fmt.Println("Employee Age    : ",emp2.age)
    fmt.Println("Employee Sex    : ",emp2.sex)
    
  }
    
    
	