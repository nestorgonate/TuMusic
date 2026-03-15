package playingsong

import (
	"log"
	"tumusic/models"
	"tumusic/playsong"

	tea "github.com/charmbracelet/bubbletea"
)

func (m *EmbedPlayer) PlaySongCmd() tea.Cmd {
	return func() tea.Msg {
		controller, err := playsong.PlaySong(m.PathSong)
		if err != nil {
			log.Print(err)
			return nil
		}
		return models.AudioStartedMsg{PlayerController: controller}
	}
}

func (m *EmbedPlayer) WatchingSongStopCmd() tea.Cmd{
	return func() tea.Msg{
		<-m.Player.PlayerController.Done
		return models.StopMsg{}
	}
}
