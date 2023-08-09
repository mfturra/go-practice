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
	_, err = db.Exec(insertData, "NFL Street", "Playstation 2", "2004-01-13", "EA Sports BIG")
	if err != nil {
		panic(err)
	}
	fmt.Println("Data insertion successful!")

	// Query Table
	// rows, err := db.Query(fmt.Sprintf("SELECT * FROM %s", table))
	// msg := "Error executing query"
	// if err != nil {
	// 	log.Fatalf("%s: %v", msg, err)
	// }
	// defer rows.Close()
}
