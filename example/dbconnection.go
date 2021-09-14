package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func insert(u User) {
	insert, err := sqldb.db.Query("INSERT INTO User VALUES ( '" + u.ID + "', '" + u.PWD + "' )")

	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
}

func selectAll() []User {
	results, err := sqldb.db.Query("SELECT * FROM User")
	if err != nil {
		panic(err.Error())
	}

	user := []User{}

	for results.Next() {
		u := new(User)

		err = results.Scan(&u.ID, &u.PWD)
		if err != nil {
			panic(err.Error())
		}
		user = append(user, *u)

	}
	return user
}

type DB struct {
	db  *sql.DB
	err error
}

var sqldb DB

func connection() {
	sqldb = DB{}
	fmt.Println("Go MySQL Tutorial")

	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/study")
	sqldb.db = db
	sqldb.err = err
	if err != nil {
		panic(err.Error())
	}

}

func close() {
	sqldb.db.Close()
}
