package main

import (
	"./crawl"
	"./db"
	"./server"
	"flag"
	"time"
)

func launchCrawl(path string) {
	dtb := db.DBManager{}
	dtb.GetConnection()
	c := make(chan db.Track)
	crawler := crawl.Crawler{c}
	go dtb.ReadMusicChan(c)
	crawler.Go(path)
	time.Sleep(10000)
	dtb.Close()
}

func launchServer() {
	server.Launch()
}

func main() {
	flag.Parse()
	fl := flag.Args()
	if len(fl) > 0 {
		launchCrawl(fl[0])
		return
	}
	launchServer()
}
