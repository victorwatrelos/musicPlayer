package main

/*
import (
	"./player"
	"flag"
	"time"
)

func main() {
	flag.Parse()
	dtb := db.DBManager{}
	dtb.GetConnection()
	c := make(chan db.Music)
	crawler := crawl.Crawler{c}
	go dtb.ReadMusicChan(c)
	//	crawler.Go(".")
	root := flag.Arg(0)
	crawler.Go(root)
	time.Sleep(10000)
	//	db.Query("artist", "O")
	dtb.Close()
}

*/
import (
	"./client"
	"./server"
)

func main() {
	//	server.Launch()
	client.Launch()
}
