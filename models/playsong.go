package models

import (
	"os/exec"

	"github.com/faiface/beep"
)

type playerState int

const (
	Play playerState = iota
	Pause
	Stop
)

type Player struct {
	Title            string
	PathSong         string
	State            playerState
	Cursor           int
	PlayerController *PlayerController
	Buttons          []string
}

type PlayerController struct {
	Ctrl     *beep.Ctrl
	Streamer beep.StreamCloser
	Cmd      *exec.Cmd
}

type AudioStartedMsg struct {
	PlayerController PlayerController
}

type Closer struct {
    beep.Streamer
    Cleanup func() error
}

func (c Closer) Close() error {
    return c.Close()
}

type StopMsg struct{}