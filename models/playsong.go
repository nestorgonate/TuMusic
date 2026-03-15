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
	Done     chan struct{}
}

type AudioStartedMsg struct {
	PlayerController *PlayerController
}

type Closer struct {
	Streamer beep.Streamer
	Cleanup  func() error
}

func (c Closer) Close() error {
	return c.Cleanup()
}

func (c Closer) Err() error {
	return c.Streamer.Err()
}

func (c Closer) Stream(samples [][2]float64) (int, bool) {
	return c.Streamer.Stream(samples)
}

type StopMsg struct{}
