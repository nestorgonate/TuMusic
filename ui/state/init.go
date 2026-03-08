package state

import tea "github.com/charmbracelet/bubbletea"

func (m mainModel) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}