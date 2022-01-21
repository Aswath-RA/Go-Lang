package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //You could import dialect
)

type User struct {
	ID       int
	Username string
}

func main() {
	db, err := gorm.Open("mysql", "blackfox:********@tcp(127.0.0.1:3306)/Employee?charset=utf8&parseTime=True")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	defer fmt.Println("Connection closed")
	fmt.Println("Connection enabled")

	db.CreateTable(&User{})
	fmt.Println("Table Created")

	//Inserting Single value

	//Inserting Multiple value
	var users []User = []User{
		User{ID: 1, Username: "Rock"},
		User{ID: 2, Username: "John"},
		User{ID: 3, Username: "Carter"},
	}

	for _, val := range users {
		db.Create(&val)

	}

	fmt.Println("Value inserted")

	var Getid int
	var Getname string
	fmt.Println("Enter the Id to change: ")
	fmt.Scan(&Getid)
	fmt.Println("Enter the Username to change: ")
	fmt.Scan(&Getname)

	// Update
	for _, val := range users {
		if val.ID == Getid {
			db.Model(&val).Update("ID", Getid)

			db.Model(&val).Updates(
				map[string]interface{}{
					"Username": Getname,
				})
		}

	}
	fmt.Println("Value Updated")

	db.HasTable(&User{})
	//Deleting table
	db.DropTableIfExists(&User{})
	fmt.Println("Table deleted")

}
