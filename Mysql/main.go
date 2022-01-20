package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//DataBase connection
	fmt.Println("Mysql")
	db, err := sql.Open("mysql", "blackfox:11aa22**@tcp(0.0.0.0:3306)/Employee")
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Database Connected")

	//Table creation

	_, err = db.Exec("CREATE TABLE Employee(id INT NOT NULL, name VARCHAR(20) NOT NULL,age INT NOT NULL,PRIMARY KEY (ID));")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Table Created")
	defer db.Close()

	//Value Insertion

	_, err = db.Exec("INSERT Employee (id,name,age) VALUES(1,'Aswath',21);")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Data Inserted")

	//Retriewing values from table

	result, err := db.Query("SELECT * FROM Employee;")

	if err != nil {
		log.Fatal(err)
	}

	for result.Next() {
		var id int
		var name string
		var age int

		err := result.Scan(&id, &name, &age)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, " ", name, " ", age)

	}

	//Updating table

	_, err = db.Exec("update Employee set name = 'sasi' where id =1;")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Value updated")

	//Output:

	//Retriewing values from  updated table

	resu, err := db.Query("SELECT * FROM Employee;") //Displaying the updated value

	if err != nil {
		log.Fatal(err)
	}

	for resu.Next() {
		var id int
		var name string
		var age int

		err := resu.Scan(&id, &name, &age)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, " ", name, " ", age)

	}

	//Deleting table
	_, err = db.Exec("Drop table Employee;")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Employee table Deleted")

}
