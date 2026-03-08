package playingsong

import (
	"log"
	"tumusic/playsong"
	tea "github.com/charmbracelet/bubbletea"
)

func (m EmbedPlayer) Init() tea.Cmd {
	return func() tea.Msg {
		err := playsong.InitSpeaker()
		if err != nil{
			log.Print(err)
		}
		return nil
	}
}
