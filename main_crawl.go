package main

import (
	"./crawl"
	"./db"
	"flag"
	"time"
)

func main() {
	flag.Parse()
	fl := flag.Arg(0)
	dtb := db.DBManager{}
	dtb.GetConnection()
	c := make(chan db.Track)
	crawler := crawl.Crawler{c}
	go dtb.ReadMusicChan(c)
	//	crawler.Go("/Users/josephrenouf/Music/iTunes/iTunes Media/Music")
	crawler.Go(fl)
	time.Sleep(10000)
	//  db.Query("artist", "O")
	dtb.Close()
}
