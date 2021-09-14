package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID  string
	PWD string
}

func main() {
	fmt.Println("Go MySQL Tutorial")

	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/study")

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// insert, err := db.Query("INSERT INTO User VALUES ( 'visitant2', '1234' )")

	// if err != nil {
	// 	panic(err.Error())
	// }
	// defer insert.Close()

	results, err := db.Query("SELECT * FROM User")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for results.Next() {
		var user User
		// for each row, scan the result into our tag composite object
		err = results.Scan(&user.ID, &user.PWD)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
		fmt.Println(user.ID, " : ", user.PWD)

	}
}
