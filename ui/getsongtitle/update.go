package getsongtitle

import tea "github.com/charmbracelet/bubbletea"

func (m EmbedGetSongTitleMenu) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			m.TextInput.Value()
			return m, nil
		case "esc":
			m.TextInput.Blur()
			return m, nil
		case "ctrl+c":
			return m, tea.Quit
		}
	}
	m.TextInput, cmd = m.TextInput.Update(msg)
	return m, cmd
}
