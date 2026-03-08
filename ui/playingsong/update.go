package playingsong

import (
	"tumusic/models"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/faiface/beep/speaker"
)

func (m EmbedPlayer) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type){
	case models.AudioStartedMsg:
		m.Player.PlayerController = &msg.PlayerController
		m.Player.State = models.Play
		return m, nil
	case tea.KeyMsg:
		switch msg.String(){
		case "up":
			if m.Player.Cursor > 0{
				m.Player.Cursor--
			}
		case "down":
			if m.Player.Cursor < 2{
				m.Player.Cursor++
			}
		case "enter":
			switch m.Player.Cursor{
			//0 pause/resume song
			case 0:
				if m.Player.PlayerController != nil {
					speaker.Lock()
					m.Player.PlayerController.Ctrl.Paused = !m.Player.PlayerController.Ctrl.Paused
					speaker.Unlock()
					if m.Player.PlayerController.Ctrl.Paused {
						m.Player.State = models.Pause
					} else {
						m.Player.State = models.Play
					}
				}
			//1 stop music
			case 1:
				if m.Player.PlayerController != nil {
					speaker.Lock()
					speaker.Clear()
					speaker.Unlock()
					_ = m.Player.PlayerController.Streamer.Close()
					m.Player.State = models.Stop
					m.Player.PlayerController = nil
				}
				return m, func() tea.Msg { return models.StopMsg{} }
			}
		case "ctrl+c":
			return m, tea.Quit
		}
	}
	return m, nil
}