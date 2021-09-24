package main

import(
    "fmt"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

// Declaring some global variables
var user = "root"
var passwd = "############"     // Make sure to put your mySQL password
var serverIP = "127.0.0.1:3306"
var dbName = "classexample1"    // This is the database name

func main(){
    fmt.Println("Go mySQL Example")

    // open up database connection
    // create the database in your local machine or wherever
    db, err := sql.Open("mysql", user+":" +passwd +"@tcp(" +serverIP +")/" +dbName)

    // Handle for error while establishing connection
    if err != nil{
        fmt.Println("Error establishing connection...")
        panic(err.Error())
    }

    // wait until everything else finishes executing and then close out the sql connection
    defer db.Close()

    // insert a sid, sname, color [Make sure to open an sql connection before]
    insertBoat(db, "303", "BlackCherry","Black")
}

// This function will allow us to insert a row into our boats table
func insertBoat(db *sql.DB, bid ,bname, color string) {
    insert, err := db.Query("INSERT INTO boats VALUES ("+bid+","+"\""+bname+"\""+","+"\""+color+"\""+")")

    if err != nil {
        fmt.Println("Error while inserting to table")
        panic(err.Error())
    }

    defer insert.Close()
}
