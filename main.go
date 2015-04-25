package main

import (
	"./crawl"
	"./db"
)

func main() {
	db := db.DBManager{}
	db.GetConnection()
	c := make(chan crawl.Music)
	db.Query("artist", "O")
	db.Close()
}
