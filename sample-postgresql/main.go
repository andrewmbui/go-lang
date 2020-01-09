package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Magic123"
	dbname   = "socratica"
)

type book struct {
	ID    string `json:"id"`
	Isbn  string `json:"isbn"`
	Title string `json:"title"`
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	 Query All Records
	sqlStatement := `
		SELECT 	uniqueid,
				isbn,
				title
		FROM	books`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		bk := book{}
		err = rows.Scan(
			&bk.ID,
			&bk.Isbn,
			&bk.Title,
		)
		if err != nil {
			panic(err)
		}
		fmt.Println("Record!")
		fmt.Println("ID: ", bk.ID, " ISBN: ", bk.Isbn, " Title: ", bk.Title)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	

	statement := `INSERT INTO books(uniqueid, isbn, title) VALUES (3, 098765, 'Book Three')`
	stmt, err := db.Query(statement)
	if err != nil {
		fmt.Print(err)
	}
	defer stmt.Close()

}
