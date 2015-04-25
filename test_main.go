package main

import (
	"./crawl"
	"flag"
)

func main() {
	flag.Parse()
	root := flag.Arg(0)
	crawl.Crawl(root)
}
