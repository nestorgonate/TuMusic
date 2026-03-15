package playsong

import (
	"fmt"
	"io"
	"os/exec"
	"path/filepath"
	"time"
	"tumusic/models"

	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
)

type EmbedPlayerController struct {
	models.PlayerController
}

var format = beep.Format{
	SampleRate:  44100,
	NumChannels: 2,
	Precision:   2,
}

func InitSpeaker() error {
	return speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/2))
}

func PlaySong(pathSong string) (*models.PlayerController, error) {
	absPath, err := filepath.Abs(pathSong)
	done := make(chan struct{})
	if err != nil {
		return &models.PlayerController{}, fmt.Errorf("filepath abs: %w", err)
	}

	cmd := exec.Command("ffmpeg", "-i", absPath, "-acodec", "pcm_s16le", "-f", "wav", "pipe:1")
	cmd.Stderr = nil

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return &models.PlayerController{}, fmt.Errorf("stdout pipe: %w", err)
	}
	if err = cmd.Start(); err != nil {
		return &models.PlayerController{}, fmt.Errorf("ffmpeg start: %w", err)
	}

	wavData, err := io.ReadAll(stdout)
	cmd.Wait()
	if err != nil {
		return &models.PlayerController{}, fmt.Errorf("leer wav: %w", err)
	}

	pcmData := wavData[44:]
	pos := 0
	streamer := beep.StreamerFunc(func(samples [][2]float64) (int, bool) {
		numSamples := len(samples)
		framesAvail := (len(pcmData) - pos) / 4
		if framesAvail <= 0 {
			return 0, false
		}
		if framesAvail < numSamples {
			numSamples = framesAvail
		}
		for i := 0; i < numSamples; i++ {
			l := int16(pcmData[pos]) | int16(pcmData[pos+1])<<8
			r := int16(pcmData[pos+2]) | int16(pcmData[pos+3])<<8
			samples[i][0] = float64(l) / 32768.0
			samples[i][1] = float64(r) / 32768.0
			pos += 4
		}
		return numSamples, true
	})

	ctrl := &beep.Ctrl{Streamer: streamer, Paused: false}
	speaker.Play(beep.Seq(ctrl, beep.Callback(func() {
		close(done)
	})))

	return &models.PlayerController{
		Ctrl: ctrl,
		Streamer: models.Closer{ // Closer implementa StreamCloser porque tiene Close()
			Streamer: streamer,
			Cleanup: func() error {
				if cmd.Process != nil {
					return cmd.Process.Kill()
				}
				return nil
			},
		},
		Cmd: cmd,
		Done: done,
	}, nil
}