package main

import (
	"./db"
	"./player"
	"flag"
)

func main() {
	bdd := db.DBManager{}
	bdd.Query("artist", "NOFX")
	flag.Parse()
	play := player.Player{}
	play.Add(flag.Arg(0))
	play.Play()
}
