package startmenu

import tea "github.com/charmbracelet/bubbletea"

func (m EmbedingStartMenu) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type){
	case tea.KeyMsg:
		switch msg.String(){
		case "up":
			if m.Cursor > 0{
				m.Cursor--
			}
		case "down":
			if m.Cursor < len(m.Choices)-1{
				m.Cursor++
			}
		case "enter":
			_, ok := m.Selected[m.Cursor]
			if ok{
				m.Selected[m.Cursor] = struct{}{}
			}
		case "ctrl+c":
			return m, tea.Quit
		}
	}
	return m, nil
}