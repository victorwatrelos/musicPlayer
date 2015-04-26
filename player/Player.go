package player

import (
	"fmt"
	"io"
	"os/exec"
)

type Player struct {
	input io.WriteCloser
	cmd   *exec.Cmd
}

func (p *Player) Init() {
	fmt.Println("Starting vlc")
	var err error
	p.cmd = exec.Command("vlc", "-I", "rc")
	p.cmd.Start()
	p.input, err = p.cmd.StdinPipe()
	chk(err)
}

func (p *Player) Play() {

	_, err := p.input.Write([]byte("play\n"))
	chk(err)
}

func (p *Player) Add(filename string) {
	cmd := "add ../music.mp3\n"
	cmd2 := []byte(cmd)
	fmt.Println(cmd2)
	fmt.Println(len(cmd2))
	_, err := p.input.Write(cmd2)
	chk(err)
}

func (p *Player) Close() {
	p.input.Close()
}

func chk(err error) {
	if err != nil {
		panic(err)
	}
}
