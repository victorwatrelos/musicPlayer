package player

import (
	"../vlcControl"
)

type Player struct {
	vlc vlcControl.VlcControl
}

func (p *Player) Init() {
	p.vlc.Init()
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

func (p *Player) ClearQueue() {
	p.vlc.ClearQueue()
}

func (p *Player) Close() {
	p.vlc.Close()
}

func (p *Player) ShowPlaylist() string {
	return p.vlc.ShowPlaylist()
}

func (p *Player) ShowStatus() string {
	return p.vlc.ShowStatus()
}

func (p *Player) Goto(index string) {
	p.vlc.Goto(index)
}

func (p *Player) SetLoop(act bool) {
	p.vlc.SetLoop(act)
}

func (p *Player) SetRandom(act bool) {
	p.vlc.SetRandom(act)
}

func (p *Player) SetRepeat(act bool) {
	p.vlc.SetRepeat(act)
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
