package crawl

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
)

func isAudio(file string) bool {
}

func visit(path string, f os.FileInfo, err error) error {
	fmt.Printf("Visited: %s\n", path)
	//	out, err := exec.Command("mediainfo", "-f", path).Output()
	out := exec.Command("mediainfo", "-f", path)
	fmt.Println("dfdf")
	ou, err := out.Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("info: %s\n", ou)
	return nil
}

func Crawl(dir string) {
	err := filepath.Walk(dir, visit)
	fmt.Printf("filepath.Walk() returned %v\n", err)
}
