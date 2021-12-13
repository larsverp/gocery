package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

var databaseURL string
var databasePort string
var databaseUser string
var databasePassword string
var databaseName string
var databaseType string
var databaseConnection bool = false
var db *sql.DB

func main() {
	fmt.Println("Welcome to Gocery, you're about to make your grocery shopping life a lot easier!")
	for !databaseConnection {

		fmt.Println("At the moment Gocery only supports MySQL databases, more are coming soon!")
		fmt.Println("Select your database: [0]MySQL")
		var databaseChoice int = 0
		fmt.Scanln(&databaseChoice)
		for databaseChoice != 0 {
			fmt.Println("Please enter a valid database type: [0]MySQL [1]Postgres")
			fmt.Scanln(&databaseChoice)
		}

		fmt.Println("Please enter a database URL: [127.0.0.1]")
		fmt.Scanln(&databaseURL)
		if databaseURL == "" {
			databaseURL = "127.0.0.1"
		}

		fmt.Println("Please enter your database port: [3306]")
		fmt.Scanln(&databasePort)
		if databasePort == "" {
			databasePort = "3306"
		}

		fmt.Println("Database username: [root]")
		fmt.Scanln(&databaseUser)
		if databaseUser == "" {
			databaseUser = "root"
		}

		fmt.Println("Database password:")
		fmt.Scanln(&databasePassword)

		fmt.Println("Database name: [gocery]")
		fmt.Scanln(&databaseName)
		if databaseName == "" {
			databaseName = "gocery"
		}

		databaseURL = databaseUser + ":" + databasePassword + "@tcp(" + databaseURL + ":" + databasePort + ")/" + databaseName

		// More database types should be added in the future
		if databaseChoice == 0 {
			databaseType = "mysql"
		}

		fmt.Println("Attempting to connect to DB...")
		var err error
		db, err = sql.Open(databaseType, databaseURL)
		if err != nil {
			fmt.Println("Error connecting to database: ", err)
			fmt.Println("Please try again.")
		} else {
			fmt.Println("Successfully connected to database!")
			databaseConnection = true
		}
	}

	startMigration(*db)

	startWebServer()
}

func startMigration(db sql.DB) {
	fmt.Println("Migrating database...")
	query, err := ioutil.ReadFile("migrations/initial.sql")
	if err != nil {
		panic(err)
	}
	if _, err := db.Exec(string(query)); err != nil {
		panic(err)
	}
	fmt.Println("Database migrated! You are ready to GO........cery ;) ")
}
