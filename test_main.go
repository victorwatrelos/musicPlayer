package main

import (
	"./crawl"
	"flag"
)

func main() {
	ch := make(chan crawl.Music)
	craw := crawl.Crawler{ch}
	flag.Parse()
	root := flag.Arg(0)
	craw.Go(root)
}
