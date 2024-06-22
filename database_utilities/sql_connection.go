package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func displayTableContent() {
	db, err := sql.Open("mysql", "root:Noncore123@tcp(localhost:3306)/newsarticles")
	if err != nil {
	panic(err.Error())
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Successfully connected to DATABASE newsarticles")
	
	tableName := "nbcnews"
	rows, err := db.Query("SELECT * FROM " + tableName)
	if err!= nil {
		panic(err.Error())
	}

	defer rows.Close()

	for rows.Next(){
		var id int
		var Title string
		var Date string
		var URL string
		var Synopsis string
		var Content string

		err := rows.Scan(&id, &Title, &Date, &URL, &Synopsis, &Content)
		if err != nil {
			panic(err.Error())
		}

		fmt.Println("ID: ", id, "Title: ", Title, "Date: ", Date, "URL: ", URL, "Synopsis: ", Synopsis, "Content: ", Content)
		

	}

	if err = rows.Err(); err != nil {
		panic(err.Error())
	}
}
