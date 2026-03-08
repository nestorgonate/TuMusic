package addsong

import (
	"tumusic/models"

	tea "github.com/charmbracelet/bubbletea"
)

func (m EmbedingAddSongMenu) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type){
	case models.SearchResultMsg:
		m.Songs = msg
	case tea.KeyMsg:
		switch msg.String(){
		case "ctrl+c":
			return m, tea.Quit
		case "up":
			if m.Cursor > 0{
				m.Cursor--
			}
		case "down":
			if m.Cursor < len(m.Songs)-1{
				m.Cursor++
			}
		case "enter":
			_, ok := m.Selected[m.Cursor]
			if ok{
				m.Selected[m.Cursor] = struct{}{}
			}
		}
	}
	return m, nil
}