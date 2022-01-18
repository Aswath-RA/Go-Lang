package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
)

func main() {
	os.Remove("/home/blackfox/Desktop/gotutor/Sqlite/Database/sqlite.db") //Removing file if already created
	fmt.Println("Creating Database")
	file, err := os.Create("/home/blackfox/Desktop/gotutor/Sqlite/Database/sqlite.db") //Creating a db file
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close() //closing db
	fmt.Println("Database Created")

	sqlitedatabase, err := sql.Open("sqlite3", "/home/blackfox/Desktop/gotutor/Sqlite/Database/sqlite.db") //opening db file
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database Opened")
	defer sqlitedatabase.Close()
	CreateTable(sqlitedatabase)
	defer fmt.Println("Database Closed")

	Inserttable(sqlitedatabase, 101, "Aswath", 21) //Passing values to be inserted
	Inserttable(sqlitedatabase, 102, "Carter", 45)
	Inserttable(sqlitedatabase, 103, "Bruce ", 31)
	Inserttable(sqlitedatabase, 104, "Red", 24)
	Inserttable(sqlitedatabase, 105, "Wick", 55)
	Inserttable(sqlitedatabase, 106, "Luther", 47)

	Displaytable(sqlitedatabase)
	Deletetable(sqlitedatabase)

}

func CreateTable(sqlite *sql.DB) { //Table Creation
	create_table_query := `CREATE TABLE Employee("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT ,
							"name" TEXT,"age" integer);` //sql statement for table creation.
	create_table, _ := sqlite.Prepare(create_table_query)
	create_table.Exec()
	fmt.Println("Table Created")
}

func Inserttable(sqlite *sql.DB, id int, name string, age int) { //Value Insertion
	insert_table_query := `INSERT INTO Employee (id,name,age) VALUES(?,?,?)`
	insertquery, err := sqlite.Prepare(insert_table_query)
	if err != nil {
		log.Fatal(err)
	}
	_, err = insertquery.Exec(id, name, age)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println("Data Inserted into Table")
}
func Deletetable(sqlite *sql.DB) { //Table Deletion
	delete_table_query := `DROP TABLE Employee`
	deletequery, err := sqlite.Prepare(delete_table_query)
	if err != nil {
		log.Fatal(err)
	}
	deletequery.Exec()
	fmt.Println("Table Deleted")
}
func Displaytable(sqlite *sql.DB) { //Display Table
	display, err := sqlite.Query("SELECT * FROM Employee ORDER BY id")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Employee Details:")
	fmt.Println()
	fmt.Println("ID  Name		Age")
	defer display.Close()
	for display.Next() {
		var id int
		var name string
		var age int
		display.Scan(&id, &name, &age)
		fmt.Println(id, " ", name, "  ", age)
	}
}
