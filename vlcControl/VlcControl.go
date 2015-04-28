package vlcControl

import (
	"fmt"
	"io"
	"os/exec"
	"strconv"
	"time"
)

type VlcControl struct {
	input  io.WriteCloser
	output io.ReadCloser
	cmd    *exec.Cmd
}

func (v *VlcControl) Init() {
	fmt.Println("VLC start")
	var err error
	v.cmd = exec.Command("/Applications/VLC.app/Contents/MacOS/VLC", "-I", "rc")
	v.input, err = v.cmd.StdinPipe()
	v.output, err = v.cmd.StdoutPipe()
	chk(err)
	errr := v.cmd.Start()
	chk(errr)
	time.Sleep(100)
}

func (v *VlcControl) RunCmd(cmd string) {
	println("Cmd: ", cmd)
	_, err := v.input.Write([]byte(cmd))
	chk(err)
}

func (v *VlcControl) Pause() {
	v.RunCmd("pause\n")
}

func (v *VlcControl) Play() {
	v.RunCmd("play\n")
}

func (v *VlcControl) Prev() {
	v.RunCmd("prev\n")
}

func (v *VlcControl) Next() {
	v.RunCmd("next\n")
}

func (v *VlcControl) SetVolume(vol int) {
	if vol > 1024 {
		vol = 1024
	}
	if vol < 0 {
		vol = 0
	}
	v.RunCmd("volume " + strconv.Itoa(vol) + "\n")
}

func (v *VlcControl) Add(filename string) {
	v.RunCmd("add " + filename + "\n")
}

func (v *VlcControl) AddToQueue(filename string) {
	v.RunCmd("enqueue " + filename + "\n")
}

func (v *VlcControl) ClearQueue() {
	v.runCmd("clear")
}

func (v *VlcControl) Close() {
	println("VLC shutdown")
	v.RunCmd("shutdown")
	v.input.Close()
}

func chk(err error) {
	if err != nil {
		panic(err)
	}
}
