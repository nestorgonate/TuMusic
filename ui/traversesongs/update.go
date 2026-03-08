package getdownloadedsongs

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m EmbedGetDownloadedSongs) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type){
	case tea.KeyMsg:
		switch msg.String(){
		case "ctrl+c":
			return m, tea.Quit
		case "up":
			if m.Cursor > 0{
				m.Cursor--
			}
		case "down":
			if m.Cursor < len(m.SongsFound)-1{
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
