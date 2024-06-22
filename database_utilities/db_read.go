package main

import (
	"database/sql"
	"log"
 	_ "github.com/go-sql-driver/mysql"
)

//Function to get the article Titles. The function will query the database for the required column
func getLinkList(db *sql.DB, tableName string) ([]string, error){
	query := "SELECT URL FROM " + tableName
	linkList := []string{}

	rows, err := db.Query(query)
	if err != nil {
		log.Println("Failed to Query the DATABASE TABLE %s", tableName)
		panic(err.Error())
		return linkList, err
	}

	for rows.Next(){
		var link string
		err := rows.Scan(&link)
		if err != nil {
			return linkList, err
		}
		linkList = append(linkList, link)
	}

	return linkList, nil
}

func getLinksFromDB(dbName string, tableName string) ([]string, error) {
	//connect to the Database
	dbCredentials := "root:Noncore123@tcp(localhost:3306)/" + dbName
	db, err := sql.Open("mysql", dbCredentials)
	if err!= nil {
		log.Println("Failed connecting to DATABASE ", err)
		panic(err.Error())
	}

	articleLinkList, err := getLinkList(db, tableName)
	if err != nil {
		return []string{}, err
	}

	return articleLinkList, nil
}
