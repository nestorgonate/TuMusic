package main

import (
	"fmt"
	"os/exec"
	"tumusic/ui/state"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/faiface/beep/speaker"
)

func main() {
	_, err := exec.LookPath("ffmpeg")
	if err != nil {
		fmt.Print("You have to install ffmpeg to use this program https://www.ffmpeg.org/")
	}
	menu := tea.NewProgram(state.InitialModel(), tea.WithAltScreen())
	if _, err := menu.Run(); err != nil {
		fmt.Printf("Failed to loading menu: %v", err)
	}
	speaker.Clear()
}
