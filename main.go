package main

import (
	"./bdd"
	"./crawl"
	"fmt"
)

func main() {
	db := bdd.BddManager{}
	db.GetConnection()
	data := crawl.Music{1212, 44100, 666, "NOFX", "PROUT", "SSSS", "My VV", "PATH"}
	db.InsertData(&data)
	db.Close()
}
