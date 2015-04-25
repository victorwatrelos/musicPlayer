package main

import (
	"./db"
)

func main() {
	db := db.DBManager{}
	db.GetConnection()
	db.Query("artist", "O")
	db.Close()
}
