package player

import (
	"code.google.com/p/portaudio-go/portaudio"
	"fmt"
	"os"
)

type Player struct {
}

func (p *Player) PlaySound(filename string) {
	fmt.Println("Start playing: ", filename)
	f, err := os.Open(filename)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	portaudio.Initialize()
	defer portaudio.Terminate()
	out := make([]int32, 8192)
	stream, err := portaudio.OpenDefaultStream(0, 2, 44100, len(out), &out)
	chk(err)
	defer stream.Close()
}
