package crawl

import (
	"../db"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	_ "strconv"
	"strings"
)

type Crawler struct {
	Channel chan db.Music
}

func (c *Crawler) isAudio(file string) bool {
	ext := path.Ext(file)
	if ext != ".mp3" && ext != ".m4a" && ext != ".flac" {
		return false
	}
	return true
}

func (c *Crawler) sendInfo(path string, info string) {
	mus := db.Music{"", 0, "", "Untitled", "Undefined", "", path}
	a1 := strings.Split(info, "\n")
	for _, v := range a1 {
		a2 := strings.Split(v, " : ")
		switch strings.Trim(a2[0], " ") {
		case "Duration":
			mus.Duration = a2[1]
		case "Channel(s)":
			//			mus.NbChan, _ = strconv.ParseInt(a2[1], 10, strconv.IntSize)
			mus.NbChan = 2
		case "Performer":
			mus.Artist = a2[1]
		case "Album/Performer":
			mus.Artist = a2[1]
		case "Album":
			mus.Album = a2[1]
		case "Genre":
			mus.Genre = a2[1]
		case "Track name":
			mus.Name = a2[1]
		}
	}
	c.Channel <- mus
	fmt.Println(mus)
}

func (c *Crawler) getInfo(path string, f os.FileInfo, err error) error {
	if c.isAudio(path) {
		out, err := exec.Command("mediainfo", path).Output()
		if err != nil {
			log.Fatal(err)
		}
		c.sendInfo(path, string(out))
	}
	return nil
}

func (c *Crawler) Go(dir string) {
	err := filepath.Walk(dir, c.getInfo)
	if err != nil {
		log.Fatal(err)
	}
	close(c.Channel)
}
