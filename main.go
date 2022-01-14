package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

var dbsettings = DatabaseSettings{}

func main() {
	fmt.Println("Welcome to Gocery, you're about to make your grocery shopping life a lot easier!")

	if _, err := os.Stat("data.json"); errors.Is(err, os.ErrNotExist) {
		startDBQuestions(&dbsettings)

		dbsettings.startMigration()

		file, _ := json.MarshalIndent(dbsettings, "", " ")

		_ = ioutil.WriteFile("data.json", file, 0644)
	} else {
		file, _ := ioutil.ReadFile("data.json")
		_ = json.Unmarshal([]byte(file), &dbsettings)
		err := dbsettings.connect()
		if err != nil {
			fmt.Println("Error connecting to database: ", err)
		}
	}

	startWebServer()
}
