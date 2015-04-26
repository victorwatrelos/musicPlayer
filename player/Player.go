package player

import (
	"fmt"
	"io"
	"os/exec"
	"time"
)

type Player struct {
	input  io.WriteCloser
	output io.ReadCloser
	cmd    *exec.Cmd
}

func (p *Player) Init() {
	fmt.Println("Starting vlc")
	var err error
	p.cmd = exec.Command("/Applications/VLC.app/Contents/MacOS/VLC", "-I", "rc")
	p.input, err = p.cmd.StdinPipe()
	p.output, err = p.cmd.StdoutPipe()
	chk(err)
	errr := p.cmd.Start()
	chk(errr)
	time.Sleep(100)
}

func (p *Player) RunCmd(cmd string) {
	println("Cmd: ", cmd)
	_, err := p.input.Write([]byte(cmd))
	chk(err)
}

func (p *Player) Pause() {
	p.RunCommand("pause\n")
}

func (p *Player) Play() {
	p.RunCommand("play\n")
}

func (p *Player) SetVolume(vol int) {
	if vol > 1024 {
		vol = 1024
	}
	if vol < 0 {
		vol = 0
	}
}

func (p *Player) Add(filename string) {
	p.RunCmd("add " + filename + "\n")
}

func (p *Player) Close() {
	println("Close")
	p.input.Close()
}

func chk(err error) {
	if err != nil {
		panic(err)
	}
}
