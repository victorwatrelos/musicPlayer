package main

import (
	"./crawl"
	"./db"
	"flag"
	"time"
)

func main() {
	flag.Parse()
	db := db.DBManager{}
	db.GetConnection()
	c := make(chan crawl.Music)
	crawler := crawl.Crawler{c}
	go db.ReadMusicChan(c)
	//	crawler.Go(".")
	root := flag.Arg(0)
	crawler.Go(root)
	time.Sleep(10000)
	//	db.Query("artist", "O")
	db.Close()
}
