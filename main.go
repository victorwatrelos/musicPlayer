package main

import (
	"./crawl"
	"./db"
	"time"
)

func main() {
	db := db.DBManager{}
	db.GetConnection()
	c := make(chan crawl.Music)
	crawler := crawl.Crawler{c}
	go db.ReadMusicChan(c)
	crawler.Go(".")
	time.Sleep(10000)
	//	db.Query("artist", "O")
	db.Close()
}
