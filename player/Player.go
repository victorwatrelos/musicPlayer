package player

import (
	"../vlcControl"
	"fmt"
)

type Player struct {
	vlc vlcControl.VlcControl
}

func (p *Player) Init() {
	fmt.Println("Player Init start")
	p.vlc.Init()
	fmt.Println("Player Init end")
}

func (p *Player) Pause() {
	p.vlc.Pause()
}

func (p *Player) Play() {
	p.vlc.Play()
}

func (p *Player) Prev() {
	p.vlc.Prev()
}

func (p *Player) Next() {
	p.vlc.Next()
}

func (p *Player) SetVolume(vol int) {
	p.vlc.SetVolume(vol)
}

func (p *Player) Add(filename string) {
	p.vlc.Add(filename)
}

func (p *Player) AddToQueue(filename string) {
	p.vlc.AddToQueue(filename)
}

func (p *Player) Close() {
	p.vlc.Close()
}

func chk(err error) {
	if err != nil {
		panic(err)
	}
}

var singleton *Player = nil

func GetSing() *Player {
	if singleton == nil {
		singleton = &Player{}
		singleton.Init()
	}
	return singleton
}
