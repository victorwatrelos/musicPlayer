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
	Channel chan db.Track
}

func (c *Crawler) isAudio(file string) bool {
	ext := path.Ext(file)
	if ext != ".mp3" && ext != ".m4a" && ext != ".flac" && ext != ".wav" {
		return false
	}
	return true
}

func (c *Crawler) sendInfo(path string, info string) {
	track := db.Track{"", "Untitled", "Untitled", "Undefined", "Untitled", path}
	a1 := strings.Split(info, "\n")

	for _, v := range a1 {
		a2 := strings.Split(v, " : ")
		switch strings.Trim(a2[0], " ") {
		case "Duration":
			track.Duration = a2[1]
		case "Performer":
			track.Artist = a2[1]
		case "Album/Performer":
			track.Artist = a2[1]
		case "Album":
			track.Album = a2[1]
		case "Genre":
			track.Genre = a2[1]
		case "Track name":
			track.Name = a2[1]
		}
	}
	c.Channel <- track
}

func (c *Crawler) getInfo(path string, f os.FileInfo, err error) error {
	fmt.Println(path)
	if c.isAudio(path) {
		out, err := exec.Command("mediainfo", path).Output()
		if err != nil {
			fmt.Fprintf(os.Stderr, path, "\n")
			return nil
		}
		c.sendInfo(path, string(out))
	}
	return nil
}

/*
func PU(path string, f os.FileInfo, err error) error {
	fmt.Println(path)
	i++
	return nil
}
*/
func (c *Crawler) Go(dir string) {
	err := filepath.Walk(dir, c.getInfo)
	//err := filepath.Walk(dir, PU)
	if err != nil {
		log.Fatal(err)
	}
	close(c.Channel)
}
