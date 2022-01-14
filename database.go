package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
)

type DatabaseSettings struct {
	URL      string
	Port     string
	User     string
	Password string
	Name     string
	Type     int
	DB       *sql.DB
}

func (d *DatabaseSettings) connect() error {

	databaseURL := d.User + ":" + d.Password + "@tcp(" + d.URL + ":" + d.Port + ")/" + d.Name

	// More database types should be added in the future
	var databaseType string
	if d.Type == 0 {
		databaseType = "mysql"
	}

	fmt.Println("Attempting to connect to DB...")

	db, err := sql.Open(databaseType, databaseURL)
	if err != nil {
		fmt.Println("Error connecting to database: ", err)
		fmt.Println("Please try again.")
		return err
	} else {
		fmt.Println("Successfully connected to database!")
		d.DB = db
	}
	return nil
}

func startDBQuestions(dbsettings *DatabaseSettings) {

	databaseConnection := false

	for !databaseConnection {
		fmt.Println("At the moment Gocery only supports MySQL databases, more are coming soon!")
		fmt.Println("Select your database: [0]MySQL")
		var databaseChoice int = 0
		fmt.Scanln(&dbsettings.Type)
		for databaseChoice != 0 {
			fmt.Println("Please enter a valid database type: [0]MySQL [1]Postgres")
			fmt.Scanln(&databaseChoice)
		}

		fmt.Println("Please enter a database URL: [127.0.0.1]")
		fmt.Scanln(&dbsettings.URL)
		if dbsettings.URL == "" {
			dbsettings.URL = "127.0.0.1"
		}

		fmt.Println("Please enter your database port: [3306]")
		fmt.Scanln(&dbsettings.Port)
		if dbsettings.Port == "" {
			dbsettings.Port = "3306"
		}

		fmt.Println("Database username: [root]")
		fmt.Scanln(&dbsettings.User)
		if dbsettings.User == "" {
			dbsettings.User = "root"
		}

		fmt.Println("Database password:")
		fmt.Scanln(&dbsettings.Password)

		fmt.Println("Database name: [gocery]")
		fmt.Scanln(&dbsettings.Name)
		if dbsettings.Name == "" {
			dbsettings.Name = "gocery"
		}

		if (dbsettings.connect()) != nil {
			fmt.Println("Please try again.")
		} else {
			databaseConnection = true
		}
	}
}

func (d *DatabaseSettings) startMigration() {
	fmt.Println("Migrating database...")
	query, err := ioutil.ReadFile("migrations/initial.sql")
	if err != nil {
		panic(err)
	}
	if _, err := d.DB.Exec(string(query)); err != nil {
		panic(err)
	}
	fmt.Println("Database migrated! You are ready to GO........cery ;) ")
}
