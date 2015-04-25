package main

import (
	"./bdd"
	"fmt"
)

func main() {
	fmt.Println("HELLO")
	db := bdd.BddManager{}
	db.GetConnection()
	data := bdd.MusicData{"NOFX"}
	db.InsertData(&data)
	db.Close()
}
