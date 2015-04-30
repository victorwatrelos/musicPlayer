package player

import (
	"../vlcControl"
)

type Player struct {
	vlc    vlcControl.VlcControl
	loop   bool
	repeat bool
	random bool
	volume int
}

func (p *Player) Init() {
	p.vlc.Init()
	p.loop = false
	p.repeat = false
	p.random = false
	p.volume = 300
	p.vlc.SetLoop(p.loop)
	p.vlc.SetRepeat(p.repeat)
	p.vlc.SetRandom(p.random)
	p.vlc.SetVolume(p.volume)
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
	p.volume = vol
	p.vlc.SetVolume(p.volume)
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

func (p *Player) ToggleLoop() {
	if p.loop {
		p.loop = false
	} else {
		p.loop = true
	}
	p.vlc.SetLoop(p.loop)
}

func (p *Player) ToggleRepeat() {
	if p.repeat {
		p.repeat = false
	} else {
		p.repeat = true
	}
	p.vlc.SetRepeat(p.repeat)
}

func (p *Player) ToggleRandom() {
	if p.random {
		p.random = false
	} else {
		p.random = true
	}
	p.vlc.SetRandom(p.random)
}

func (p *Player) GetLoop() bool   { return p.loop }
func (p *Player) GetRepeat() bool { return p.repeat }
func (p *Player) GetRandom() bool { return p.random }

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
