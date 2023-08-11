package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var db *sql.DB

const (
	host = "localhost"
	port = 5432
	user = "postgres"
	// password = ""
	dbname = "golang_practice"
	table  = "videogames"
)

func main() {
	// Connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", host, port, user, dbname)

	// Open database
	var err error
	db, err = sql.Open("postgres", psqlconn)
	// CheckError("Error opening database:", err)
	if err != nil {
		fmt.Println("Error opening database:")
		return
	}

	// Close database
	defer db.Close()

	// Check DB
	err = db.Ping()
	if err != nil {
		fmt.Println("Error pinging database:")
		return
	}

	fmt.Println("Connected!")

	// Additional debug statements
	//fmt.Println("Database connection details:", psqlconn)

	// Dynamic Insertion of Data
	insertData := `INSERT INTO "videogames" ("videogame_title", "videogame_platform", "videogame_releasedate", "videogame_publisher") VALUES($1, $2, $3, $4)`
	// var videogameID uint32
	_, err = db.Exec(insertData, "Spider-Man 2", "Playstation 2", "June 28, 2004", "Activision")
	if err != nil {
		panic(err)
	}
	fmt.Println("Data insertion successful!")

	// Query Table
	rows, err := db.Query(fmt.Sprintf(`SELECT "videogame_title", "videogame_platform" FROM %s`, table))
	if err != nil {
		fmt.Println("Error executing query")
		return
	}
	defer rows.Close()

	for rows.Next() {
		var videogame_title string
		var videogame_platform string

		err = rows.Scan(&videogame_title, &videogame_platform)
		if err != nil {
			fmt.Println("Error executing query")
			return
		}

		fmt.Println(videogame_title, videogame_platform)
	}
}
