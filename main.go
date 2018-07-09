package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type UserData struct {
	ID              int
	CitizenID       string
	Firstname       string
	Lastname        string
	BirthYear       int
	FirstnameFather string
	LastnameFather  string
	FirstnameMother string
	LastnameMother  string
	SoldierID       int
	AddressID       int
}

func main() {
	db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/testsck")
	defer db.Close()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Connected success!")
	results, _ := db.Query("SELECT * FROM user")
	for results.Next() {
		var userData UserData
		err := results.Scan(
			&userData.ID,
			&userData.CitizenID,
			&userData.Firstname,
			&userData.Lastname,
			&userData.BirthYear,
			&userData.FirstnameFather,
			&userData.LastnameFather,
			&userData.FirstnameMother,
			&userData.LastnameMother,
			&userData.SoldierID,
			&userData.AddressID,
		)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(userData)
	}
}
