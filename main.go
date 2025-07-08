package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func main() {
	// connection string to database for book_storage
	connStr := "user=postgres password=star1990 dbname=book_storage sslmode=disable"

	// opening the database for owner postgres
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)

	}
	defer db.Close()

	fmt.Println("You have connected to the database successfully!")

	//	inserting a new book into database
	/*	_, err = db.Exec("INSERT INTO users (title, author, year) VALUES ($1, $2, $3)",
		"SQL & NoSQL Databases: Models, Languages, Consistency Options and Architectures for Big Data Management",
		"Andreas Meier, Michael Kaufmann",
		2019)
		if err != nil {
			panic(err)
		}

		fmt.Println("Book successfully added to database!")
	*/




	//	listing all books
	/*
	rows, err := db.Query("SELECT id, title, author, year FROM users")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var title string
		var author string
		var year int

		err := rows.Scan(&id, &title, &author, &year)
		if err != nil {
			panic(err)
		}
		fmt.Printf("ID: %d, Title: %s, Author: %s, Year: %d\n", id, title, author, year)
	}
	fmt.Println("All books successfully listed!")
	*/



	//	updating a books title

/*	_, err = db.Exec("UPDATE users SET title = $1 WHERE id = $2", "Database System Concepts", 1)
	if err != nil {
		panic(err)
	}
	fmt.Println("Book title successfully updated!")

	*/



	//	deleting book by ID

/*	_, err = db.Exec("DELETE FROM users WHERE id = $1", 1)
	if err != nil {
		panic(err)
	}
		*/
}
