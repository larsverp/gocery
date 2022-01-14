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
var user = User{}

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

	if _, err := os.Stat("userdata.json"); errors.Is(err, os.ErrNotExist) {
		user.StartQuestions()
		user.Login()

		file, _ := json.MarshalIndent(user, "", " ")

		_ = ioutil.WriteFile("userdata.json", file, 0644)
	} else {
		file, _ := ioutil.ReadFile("userdata.json")
		_ = json.Unmarshal([]byte(file), &user)
		user.Login()
	}

	startWebServer()
}
