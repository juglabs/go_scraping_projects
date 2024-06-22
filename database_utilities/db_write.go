package main 

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func writeToTable(articleList []Article) error {
	db, err := sql.Open("mysql", "root:Noncore123@tcp(localhost:3306)/newsarticles")
	if err != nil {
		log.Println("Failed to connect to DATABASE - nbcnews")
		panic(err.Error())
		return err
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Println("Failed to Ping DATABASE - nbcnews")
		panic(err.Error())
		return err
	}

	stmnt, err := db.Prepare("INSERT INTO nbcnews(Title, Date, URL, Synopsis, Content) VALUES(?, ?, ?, ?, ?)")
	if err != nil {
		log.Println("Error Preparing DATABASE")
		panic(err.Error())
		return err
	}

	for _, article := range articleList {
		_, err := stmnt.Exec(article.Title, article.Date, article.URL, article.Synopsis, article.Content)
		if err != nil {
			log.Println("Error Executing MYSQL statement ", err)
			panic(err.Error())
			return err
		}
	}

	fmt.Println("Successfully wrote all Articles to TABLE nbcnews")
	return nil
}
